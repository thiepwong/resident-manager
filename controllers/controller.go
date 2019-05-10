package controllers

import (
	"time"

	"github.com/kataras/iris"
	"github.com/thiepwong/resident-manager/services"
)

type MvcResult interface {
	GenerateResult(int, string, interface{}) (result *mvcResult)
}

type Controller struct {
	Ctx     iris.Context
	Service services.EmployeeService
	Result  MvcResult
	Auth    bool
}

type mvcResult struct {
	System      string
	Version     string
	RequestTime int64
	Code        int
	Message     string
	Data        interface{}
}

func NewMvcResult(result *interface{}) MvcResult {
	return &mvcResult{Data: result, Code: 100, Version: "1.0", System: "SmartID"}

}

func (c *mvcResult) GenerateResult(code int, msg string, d interface{}) (result *mvcResult) {

	if code == 0 {
		code = 200
	}

	if msg == "" {
		msg = "Successful"
	}

	c.RequestTime = time.Now().Unix()
	c.Code = code
	c.Message = msg
	c.Data = d
	return c
}

// func jwt(ctx iris.Context) {
// 	_auth := ctx.GetHeader("Authentication")

// }