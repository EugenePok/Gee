package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Context struct {
	Writer     http.ResponseWriter
	Req        *http.Request
	Path       string
	Method     string
	Params     map[string]string
	StatusCode int
}

func NewContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) Param(key string) string {
	val, _ := c.Params[key]
	return val
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) String(statusCode int, format string, values ...any) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(statusCode)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(statusCode int, obj any) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(statusCode)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) Data(statusCode int, data []byte) {
	c.Status(statusCode)
	c.Writer.Write(data)
}

func (c *Context) HTML(statusCode int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(statusCode)
	c.Writer.Write([]byte(html))
}
