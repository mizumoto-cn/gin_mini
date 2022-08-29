package main

import (
	"fmt"
	"net/http"

	"github.com/mizumoto-cn/ginmini/ginmini"
)

func main() {
	e := ginmini.New()
	e.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})

	e.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	e.RUN(":10086")
}
