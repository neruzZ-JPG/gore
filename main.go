package main

import (
	"gore/src/entity"
	"gore/src/route"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {

	entity.InitData()

	app := iris.New()
	mvc.Configure(app.Party("/gore"), route.MyMVC)
	app.Run(iris.Addr(":8000"))
}
