package main

import (
	"mySimpleWeb/simpleWeb"
	"net/http"
)

func main() {
	s := simpleWeb.New()
	s.GET("/", func(c *simpleWeb.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	s.GET("/hello", func(c *simpleWeb.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	s.POST("/login", func(c *simpleWeb.Context) {
		c.JSON(http.StatusOK, simpleWeb.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	s.Run(":9999")
}