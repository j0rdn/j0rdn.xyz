package main

import (
    "github.com/kataras/iris"
    "github.com/iris-contrib/template/markdown"
)

func main(){
    app := iris.New()
    app.Config.Gzip = true
    app.UseTemplate(markdown.New()).Directory("public/templates",".md")
    app.StaticFS("/public","./public/",1)
    app.Get("/", home)
    app.Listen(":80")
}


func home(ctx *iris.Context){
    ctx.MustRender("index.md",nil)
}
