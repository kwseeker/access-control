package middleware

type UrlInfo struct {
	Url    string
	Method string
}

// CasbinExclude casbin 排除的路由列表
var CasbinExclude = []UrlInfo{
	//TODO
	{Url: "/api/v1/dict/type-option-select", Method: "GET"},
	{Url: "/api/v1/dict-data/option-select", Method: "GET"},
}
