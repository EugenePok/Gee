package gee

import (
	"net/http"
)

type H map[string]interface{}

type HandlerFunc func(*Context)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (e *Engine) addRoute(method string, pattern string, handlerFunc HandlerFunc) {
	e.router.addRoute(method, pattern, handlerFunc)
}

func (e *Engine) GET(pattern string, handlerFunc HandlerFunc) {
	e.addRoute("GET", pattern, handlerFunc)
}

func (e *Engine) POST(pattern string, handlerFunc HandlerFunc) {
	e.addRoute("POST", pattern, handlerFunc)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := NewContext(w, req)
	e.router.handle(c)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}
