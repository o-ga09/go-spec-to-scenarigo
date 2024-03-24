## spec2scenarigo

## 概要／Overview

本CLIツールは、OpenAPIからScenarigoのテストシナリオを生成するためのツールです。

## インストール／Install

```bash
$ go install github.com/o-ga09/go-spec-to-scenarigo@latest
```

## 使い方／Usage

- ```csv-file```：OpenAPI Specのパラメータを指定したくない場合に、パラメーターを別途指定するcsvファイル
- ```dry-run```：シナリオファイルの生成なしに、デバッグ可能にする
- ```host```：APIサーバのURL(OpenAPI Specに複数指定されている場合に、一番最初の要素のURLが使用される仕様のため)
- ```output-file```：シナリオのファイル名を指定可能

```bash
Usage:
  go-spec-to-scenarigo [input file] [flags]

Flags:
  -c, --csv-file string      add test pattern parameter
  -d, --dry-run              dry run mode. not generate scenario file
  -h, --help                 Help message
  -s, --host string          API EndPoint
  -o, --output-file string   output file name
```

## 特徴／Features

- テスト対象のAPIのレスポンスを使用して、シナリオを生成する
- パスパラメーターや、クエリパラメーターにOpenAPIの記述を使用したくない場合にcsvで別途指定できる
- ```--dry-run```オプションを指定すると、シナリオファイルの生成なしにOpenAPI Specから値が取得できているかを確認できる

## 注意事項／Caution

- 本ツールは、開発中のため、バグが含まれているかの性があります。また、テストも全て書かれていないので品質については保証いたしかねます。

## 今後の予定／Updates

- [ ] テストを書く
- [ ] APIをリクエストする認証に対応する（basic認証、Bearer認証、awssigv4）
