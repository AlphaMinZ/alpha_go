package alpha

import "testing"

func Login(ctx *Context) {
	ctx.response.Write([]byte("login请求成功"))
}

func Register(ctx *Context) {
	ctx.response.Write([]byte("register请求成功"))
}

func TestHTTP_Start(t *testing.T) {
	h := NewHTTP()
	h.GET("/login", Login)
	h.POST("/register", Register)
	err := h.Start(":8080")
	if err != nil {
		panic(err)
	}
}
