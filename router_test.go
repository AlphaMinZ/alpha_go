package alpha

import "testing"

func TestRouterAdd(t *testing.T) {
	testCases := []struct {
		name    string
		method  string
		pattern string
	}{
		{
			name:    "test1",
			method:  "GET",
			pattern: "/study/:golang",
		},
		// {
		// 	name:    "test2",
		// 	method:  "GET",
		// 	pattern: "study",
		// },
	}
	r := newRouter()
	var mockHandleFunc HandleFunc = func(ctx *Context) {

	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r.addRouter(tc.method, tc.pattern, mockHandleFunc)
		})
	}
}
