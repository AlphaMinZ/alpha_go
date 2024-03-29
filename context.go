package alpha

import "net/http"

// Context 上下文
type Context struct {
	// 响应
	response http.ResponseWriter
	// 请求
	request *http.Request

	// 当前请求方法
	Method string
	// 请求 URL
	Pattern string
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		response: w,
		request:  r,
		Method:   r.Method,
		Pattern:  r.URL.Path,
	}
}
