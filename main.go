package main

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	app := iris.New()
	mvc.Configure(app.Party("/root"), myMVC)
	app.Run(iris.Addr(":8000"))
}

func myMVC(app *mvc.Application) {
	// app.Register(...)
	// app.Router.Use/UseGlobal/Done(...)
	app.Handle(new(MyController))
}

type MyController struct{}

func (m *MyController) BeforeActivation(b mvc.BeforeActivation) {
	// b.Dependencies().Add/Remove
	// b.Router().Use/UseGlobal/Done // 和你已知的任何标准 API  调用

	// 1-> 方法
	// 2-> 路径
	// 3-> 控制器函数的名称将被解析未一个处理程序 [ handler ]
	// 4-> 任何应该在 MyCustomHandler 之前运行的处理程序[ handlers ]
	// b.Handle("GET", "/something/{id:long}", "MyCustomHandler", anyMiddleware...)
	b.Handle("GET", "/something/{id:long}", "MyCustomHandler")
}

// GET: http://localhost:8080/root
func (m *MyController) Get() string { return "Hey" }

// GET: http://localhost:8080/root/something/{id:long}
func (m *MyController) MyCustomHandler(id int64) string { return "MyCustomHandler says Hey" }

func Middleware1(ctx iris.Context) {
	fmt.Println("hello from middleware1")
}
