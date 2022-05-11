package config

import (
	"encoding/json"
	"fmt"
)

type FormatResponse struct {
	Data    interface{}
	Code    int
	Success bool
}

func SucessResponse(data interface{}) []byte {
	formatResponse := &FormatResponse{
		Data:    data,
		Code:    CODESUCCESS,
		Success: true,
	}
	res, err := json.Marshal(formatResponse)
	if err != nil {
		fmt.Print("[successResponse] error when json")
		return nil
	}
	return res
}

func ErrorResponse(data interface{}, code int) []byte {
	formatResponse := &FormatResponse{
		Data:    data,
		Code:    code,
		Success: false,
	}
	res, err := json.Marshal(formatResponse)
	if err != nil {
		fmt.Print("[errorResoponse]error when json")
		return nil
	}
	return res
}
