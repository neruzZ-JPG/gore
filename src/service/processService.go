package service

import (
	"gore/src/config"
	"gore/src/entity"
	"gore/src/model"
)

func SelectProcesses() []byte {
	processList := model.SelectProcesses()
	res := config.SucessResponse(processList)
	return res
}

func RefreshFile() []byte {
	if err := entity.SaveFiles(); err != nil {
		res := config.ErrorResponse(nil, config.FAIL_INTERNAL)
		return res
	}
	res := config.SucessResponse(nil)
	return res
}

func RegisterProcess(name string) []byte {
	if err := model.InsertProcess(name); err != nil {
		return config.ErrorResponse(nil, config.FAIL_INTERNAL)
	}
	return config.SucessResponse(nil)
}

func DeleteProcess(name string) []byte {
	if err := model.DeleteProcess(name); err != nil {
		return config.ErrorResponse(nil, config.FAIL_INTERNAL)
	}
	return config.SucessResponse(nil)
}
