package cmd

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/getkin/kin-openapi/openapi3"
	yaml "gopkg.in/yaml.v2"
)

func GenItem(inputFileName string) (*APISpec, error) {
	// パス毎の構造体を格納するスライスを定義
	var apiSpec APISpec
	var pathSpecs []pathSpec

	// OpenAPIのYAMLファイルを読み込み
	doc, err := openapi3.NewLoader().LoadFromFile(inputFileName)
	if err != nil {
		fmt.Println("Error:", err)
		return &APISpec{}, err
	}

	// API Specのメタデータを格納
	apiSpec.Title = doc.Info.Title
	apiSpec.Description = doc.Info.Description
	apiSpec.Version = doc.Info.Version
	apiSpec.BaseUrl, _ = doc.Servers.BasePath()

	// パス毎に処理
	for _, path := range doc.Paths.InMatchingOrder() {

		// それぞれのパスに対するメソッドの一覧を取得
		obj := doc.Paths.Find(path).Operations()

		// メソッド毎の構造体を格納するスライスを定義
		var baseSpecs []baseSpec

		// メソッド毎に処理
		for method, op := range obj {

			// クエリとボディに当たるパラメータ構造体を格納するスライスを定義
			var queries []paramSpec
			var bodies []paramSpec
			var responses []responseSpec

			// 元データにクエリパラメータがある場合
			if op.Parameters != nil {
				for _, q := range op.Parameters {

					// クエリ毎にクエリパラメータ構造体を生成
					queries = append(queries, paramSpec{
						// フィールドの名前
						Name: q.Value.Name,
						// フィールドの型
						Type: q.Value.Schema.Value.Type,
						// フィールドのサンプル値
						Example: q.Value.Example,
					})
				}
			}

			// 元データにボディパラメータがある場合
			if op.RequestBody != nil {
				for name, b := range op.RequestBody.Value.Content["application/json"].Schema.Value.Properties {

					// ボディ毎にボディパラメータ構造体を生成
					bodies = append(bodies, paramSpec{
						// フィールドの名前
						Name: name,
						// フィールドの型
						Type: b.Value.Type,
						// フィールドのサンプル値
						Example: b.Value.Example,
					})
				}
			}

			if op.Responses != nil {
				for name, r := range op.Responses.Map() {
					responses = append(responses, responseSpec{
						Name:        name,
						Description: *r.Value.Description,
						Example:     r.Value.Content["application/json"].Example,
					})
				}
			}

			// メソッド毎にメソッド構造体を生成して末尾に追加
			baseSpecs = append(baseSpecs, baseSpec{
				// メソッド名
				Method: method,
				// APIのサマリ
				Summary: op.Summary,
				// ボディパラメータ構造体のスライス
				Body: bodies,
				// クエリパラメータ構造体のスライス
				Params: queries,
				// レスポンス構造体のスライス
				Response: responses,
			})

		}

		// パス毎にパス構造体を生成して末尾に追加
		pathSpecs = append(pathSpecs, pathSpec{
			// パス
			Path: path,
			// メソッド構造体のスライス
			Methods: baseSpecs,
		})

	}

	apiSpec.PathSpec = pathSpecs
	return &apiSpec, nil
}

func GenScenario(apiSpec *APISpec) error {
	// シナリオの構造体
	var scenario Scenario

	// 各ステップの構造体
	var step step

	// メタデータの構造体
	var requestInfo requestInfo
	var expectInfo expectInfo

	// シナリオを作成
	scenario.Title = apiSpec.Title

	// ステップ毎にシナリオを作成
	for _, spec := range apiSpec.PathSpec {
		for _, method := range spec.Methods {
			step.Title = method.Summary
			step.Protocol = "http"
			requestInfo.Method = method.Method
			requestInfo.Url = "https://example.com/v1" + spec.Path

			for _, r := range method.Response {
				i, _ := strconv.Atoi(r.Name)
				expectInfo.StatusCode = i
				expectInfo.Body = r.Example
				step.Request = requestInfo
				step.Expect = expectInfo
				scenario.Step = append(scenario.Step, step)
			}
		}
	}

	// シナリオファイルを作成
	f, err := os.Create("output.yaml")
	if err != nil {
		return errors.New("Scenario file cannot create")
	}
	defer f.Close()

	// シナリオをyaml形式の変換
	b, err := yaml.Marshal(scenario)
	if err != nil {
		return errors.New("Scenario file cannot convert")
	}

	// シナリオをファイルに書き込み
	_, err = f.Write(b)
	if err != nil {
		return errors.New("Scenario file cannot write")
	}

	return nil
}
