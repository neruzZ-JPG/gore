package controller

import (
	"fmt"
	"gore/src/service"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type ProcessController struct{}

//routes
func (pc *ProcessController) BeforeActivition(b mvc.BeforeActivation) {
	b.Handle("GET", "/processes", "GetProcesses")
	//b.Handle("POST", "/refreshFile", "Post")
	//b.Handle("POST", "/registerProcess", "PostRegisterProcess")
}

func (pc *ProcessController) GetProcesses() []byte {
	return service.SelectProcesses()
}

func (pc *ProcessController) PostRefresh_file(ctx iris.Context) []byte {
	return service.RefreshFile()
}

func (pc *ProcessController) PostRegister_process(ctx iris.Context) []byte {
	var jsonMap map[string]string
	_ = ctx.ReadJSON(&jsonMap)
	name := jsonMap["name"]
	fmt.Println("[registerProcess] input process name is ", name)
	return service.RegisterProcess(name)
}

func (pc *ProcessController) PostDelete_process(ctx iris.Context) []byte {
	var jsonMap map[string]string
	_ = ctx.ReadJSON(&jsonMap)
	name := jsonMap["name"]
	fmt.Println("[deleteProcess] to delete process is ", name)
	return service.DeleteProcess(name)
}
