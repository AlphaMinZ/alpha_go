package alpha

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// HandleFunc 视图函数
type HandleFunc func(ctx *Context)

type server interface {
	http.Handler

	Start(addr string) error

	Stop() error

	// addRouter
	addRouter(method string, pattern string, handleFunc HandleFunc)
}

type HTTPOption func(h *HTTPServer)

type HTTPServer struct {
	srv  *http.Server
	stop func() error

	// routers 临时存放
	routers map[string]HandleFunc
}

func WithHTTPServerStop(fn func() error) HTTPOption {
	return func(h *HTTPServer) {
		if fn == nil {
			fn = func() error {
				fmt.Println("---------------")
				quit := make(chan os.Signal)
				signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
				<-quit
				log.Println("Shutdown Server ...")

				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()
				if err := h.srv.Shutdown(ctx); err != nil {
					log.Fatal("Server Shutdown:", err)
				}
				select {
				case <-ctx.Done():
					log.Println("timeout of 5 seconds.")
				}
				return nil
			}
		}
		h.stop = fn
	}
}

func NewHTTP(opts ...HTTPOption) *HTTPServer {
	h := &HTTPServer{
		routers: map[string]HandleFunc{},
	}
	for _, opt := range opts {
		opt(h)
	}
	return h
}

// ServeHTTP 接受请求，转发请求
// 接受前端传过来的请求
// 转发前端的请求到框架
func (h *HTTPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := fmt.Sprintf("%s-%s", r.Method, r.URL.Path)
	handler, ok := h.routers[key]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("404 NOT FOUND"))
		return
	}

	// 构造当前请求的上下文
	c := NewContext(w, r)
	fmt.Printf("request %s - %s\n", c.Method, c.Pattern)
	// 转发请求
	handler(c)
}

func (h *HTTPServer) Start(addr string) error {
	h.srv = &http.Server{
		Addr:    addr,
		Handler: h,
	}
	return h.srv.ListenAndServe()
}

func (h *HTTPServer) Stop() error {
	return h.stop()
}

func (h *HTTPServer) addRouter(method string, pattern string, handleFunc HandleFunc) {
	// key is only one
	key := fmt.Sprintf("%s-%s", method, pattern)
	fmt.Printf("add router %s - %s\n", method, pattern)
	h.routers[key] = handleFunc
}

// GET get 请求
func (h *HTTPServer) GET(pattern string, handleFunc HandleFunc) {
	h.addRouter(http.MethodGet, pattern, handleFunc)
}

func (h *HTTPServer) POST(pattern string, handleFunc HandleFunc) {
	h.addRouter(http.MethodPost, pattern, handleFunc)
}

func (h *HTTPServer) DELETE(pattern string, handleFunc HandleFunc) {
	h.addRouter(http.MethodDelete, pattern, handleFunc)
}

func (h *HTTPServer) PUT(pattern string, handleFunc HandleFunc) {
	h.addRouter(http.MethodPut, pattern, handleFunc)
}
