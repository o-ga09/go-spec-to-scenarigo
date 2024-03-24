package cmd

// API Spec 構造体
type APISpec struct {
	Title       string     `yaml:"title,omitempty"`
	Description string     `yaml:"description,omitempty"`
	Version     string     `yaml:"version,omitempty"`
	BaseUrl     string     `yaml:"base_url,omitempty"`
	PathSpec    []pathSpec `yaml:"path_spec,omitempty"`
}

type pathSpec struct {
	Path    string     `yaml:"path,omitempty"`
	Methods []baseSpec `yaml:"methods,omitempty"`
}

type baseSpec struct {
	Method   string         `yaml:"method,omitempty"`
	Summary  string         `yaml:"summary,omitempty"`
	Params   []paramSpec    `yaml:"params,omitempty"`
	Body     []paramSpec    `yaml:"body,omitempty"`
	Response []responseSpec `yaml:"responses,omitempty"`
}

type paramSpec struct {
	Name    string `yaml:"name,omitempty"`
	Type    string `yaml:"param,omitempty"`
	Example any
}

type responseSpec struct {
	Name        string `yaml:"name,omitempty"`
	Description string `yaml:"description,omitempty"`
	Example     any    `yaml:"example,omitempty"`
}

// Scenarigo 構造体
type Scenario struct {
	Title string `yaml:"title,omitempty"`
	Step  []step `yaml:"steps,omitempty"`
}

type step struct {
	Title    string      `yaml:"title,omitempty"`
	Protocol string      `yaml:"protocol,omitempty"`
	Request  requestInfo `yaml:"request,omitempty"`
	Expect   expectInfo  `yaml:"expect,omitempty"`
}

type requestInfo struct {
	Method string `yaml:"method,omitempty"`
	Url    string `yaml:"url,omitempty"`
	Query  any    `yaml:"query,omitempty"`
}

type expectInfo struct {
	StatusCode int `yaml:"code,omitempty"`
	Body       any `yaml:"body,omitempty"`
}

// 追加のパラメーターの構造体
type addParam struct {
	Method string
	Query  any
	Body   string
	Path   string
}
