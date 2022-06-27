package context

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Msa map[string]any

type Context struct {
	// raw
	Writer http.ResponseWriter
	Req    *http.Request
	// request
	Path   string
	Method string
	// response
	StatusCode int
}

// NewContext returns a new Context.
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

func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) SetStatusCode(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) JSON(code int, data interface{}) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.SetStatusCode(code)
	// json.NewEncoder(c.Writer).Encode(data)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(data); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) HTML(code int, html string) {
	c.Writer.Header().Set("Content-Type", "text/html")
	c.SetStatusCode(code)
	c.Writer.Write([]byte(html))
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.Writer.Header().Set("Content-Type", "text/plain")
	c.SetStatusCode(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) Data(code int, data []byte) {
	// c.Writer.Header().Set("Content-Type", "application/octet-stream")
	c.SetStatusCode(code)
	c.Writer.Write(data)
}
