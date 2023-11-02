package main

import "github.com/kataras/iris/v12"

func main() {
	app := newApp()

	app.Listen(":8080")
}

func newApp() *iris.Application {
	app := iris.New()
	app.Get("/", index)

	return app
}

func index(ctx iris.Context) {
	accept := ctx.Request().Header.Get("Accept")
	switch accept {
	case "text/html":
		ctx.HTML("<h1>Hello World</h1>")
	case "application/json":
		var result struct {
			Text string `json:"text"`
		}
		result.Text = "Hello World"
		ctx.JSON(result)
	default:
		ctx.Text("Hello World")
	}
}
