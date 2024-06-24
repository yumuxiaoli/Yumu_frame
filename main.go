package main

import (
	"net/http"

	"github.com/yumuxiaoli/Yumu_frame/yumu"
)

func main() {
	r := yumu.New()
	r.GET("/", func(c *yumu.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Yumu</h1>")
	})
	r.GET("/hello", func(c *yumu.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	r.GET("/hello/:name", func(c *yumu.Context) {
		c.String(http.StatusOK, "hello %s,you're at %s\n", c.Param("name"), c.Path)
	})
	r.POST("/login", func(c *yumu.Context) {
		c.JSON(http.StatusOK, yumu.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	r.Run(":8888")

}
