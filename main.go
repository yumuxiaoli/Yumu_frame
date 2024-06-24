package main

import (
	"net/http"

	"github.com/yumuxiaoli/Yumu_frame/yumu"
)

func main() {
	r := yumu.New()
	r.GET("/index", func(ctx *yumu.Context) {
		ctx.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(ctx *yumu.Context) {
			ctx.HTML(http.StatusOK, "<h1>Hello Yumu</h1>")
		})

		v1.GET("/hello", func(ctx *yumu.Context) {
			ctx.String(http.StatusOK, "hello %s,you're at %s\n", ctx.Query("name"), ctx.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(ctx *yumu.Context) {
			ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Param("name"), ctx.Path)
		})
		v2.POST("/login", func(ctx *yumu.Context) {
			ctx.JSON(http.StatusOK, yumu.H{
				"username": ctx.PostForm("username"),
				"password": ctx.PostForm("password"),
			})
		})
	}
	r.Run(":8888")

}
