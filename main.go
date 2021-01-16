package main

import (
	"mySimpleWeb/simpleWeb"
	"net/http"
)

func main() {
	r := simpleWeb.Default()
	r.GET("/", func(c *simpleWeb.Context) {
		c.String(http.StatusOK, "Hello SimpleWeb\n")
	})
	r.GET("/panic", func(c *simpleWeb.Context) {
		names := []string{"SimpleWeb"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
