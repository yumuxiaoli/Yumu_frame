package main

import (
	"log"
	"net/http"
	"time"

	"github.com/yumuxiaoli/Yumu_frame/yumu"
)

func onlyForV2() yumu.HandleFunc {
	return func(ctx *yumu.Context) {
		t := time.Now()
		// ctx.Fail(500, "Internal Server Error")
		log.Printf("[%d] %s in %v for group v2", ctx.StatusCode, ctx.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := yumu.New()
	r.Use(yumu.Logger())
	r.GET("/", func(ctx *yumu.Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello Yumu</h1>")
	})
	v2 := r.Group("/v2")
	v2.Use(onlyForV2())
	{
		v2.GET("/hello/:name", func(ctx *yumu.Context) {
			ctx.String(http.StatusOK, "hello %s,you're at %s\n", ctx.Param("name"), ctx.Path)
		})
	}
	r.Run(":8888")

}
