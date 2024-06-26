package alpha

import "testing"

func Login(ctx *Context) {
	ctx.response.Write([]byte("login请求成功"))
}

func Register(ctx *Context) {
	ctx.response.Write([]byte("register请求成功"))
}

func Index(ctx *Context) {
	ctx.response.Write([]byte("首页"))
}

func ParamIndex(ctx *Context) {
	ctx.response.Write([]byte("参数路由" + ctx.params["course"]))
}

func TestHTTP_Start(t *testing.T) {
	h := NewHTTP()
	h.GET("/login", Login)
	h.GET("/study/:course", ParamIndex)
	h.POST("/register", Register)
	h.DELETE("/", Index)
	err := h.Start(":8080")
	if err != nil {
		panic(err)
	}
}
