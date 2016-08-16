import (
    "github.com/kataras/iris/"
)

func main(){
    app := iris.New()
    app.Config.Gzip = true
    app.UseTemplate(markdown.New()).Directory("./templates",".md")
    app.StaticFS("/public","./public/",1)
    app.Get("/", home)
    app.Listen(":8080")
}


func home(ctx *iris.Context){
    ctx.MustRender("index.md",nil)
}
