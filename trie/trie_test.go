package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRouter_AddRouter(t *testing.T) {
	testCases := []struct {
		name       string
		pattern    string
		data       string
		wantRouter *Router
	}{
		{
			name:    "xxx",
			pattern: "/user/login",
			data:    "hello",
			wantRouter: &Router{map[string]*Node{
				"/": {
					part: "/",
					children: map[string]*Node{
						"user": {
							part: "user",
							children: map[string]*Node{
								"login": {
									part: "login",
									data: "hello",
								},
							},
						},
					},
				},
			}},
		},
	}

	router := &Router{map[string]*Node{
		"/": {
			part: "/",
		},
	}}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			router.AddRouter(tc.pattern, tc.data)
			assert.Equal(t, tc.wantRouter, router)
		})
	}
}

func TestRouter_GetRouter(t *testing.T) {
	testCases := []struct {
		name        string
		findPattern string
		wantData    string
		wantErr     error
	}{
		{
			name:        "success",
			findPattern: "/user/login",
			wantData:    "hello",
		},
	}

	router := &Router{map[string]*Node{
		"/": {
			part: "/",
		},
	}}
	router.AddRouter("/user/login", "hello")
	router.AddRouter("/user/register", "world")
	router.AddRouter("/studry/golang", "Good")
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			n, err := router.GetRouter(tc.findPattern)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantData, n.data)
		})
	}
}
