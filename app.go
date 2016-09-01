package main

import (
    "github.com/kataras/iris"
    "github.com/iris-contrib/middleware/logger"
)

func main(){
    app := iris.New()
    app.Config.Gzip = true
    app.StaticFS("/public","/public",1)
    app.Use(logger.New(iris.Logger))
    app.OnError(iris.StatusNotFound, notFound)
    app.Get("/", home)
    app.Listen(":9000")
}

func home(ctx *iris.Context){
    ctx.MustRender("index.md",nil)
}

func notFound(ctx *iris.Context){
    ctx.MustRender("404.md",nil)
}

