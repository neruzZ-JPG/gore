package route

import (
	"gore/src/controller"

	"github.com/kataras/iris/v12/mvc"
)

func MyMVC(app *mvc.Application) {
	// app.Register(...)
	// app.Router.Use/UseGlobal/Done(...)
	app.Handle(new(controller.ProcessController))
}
