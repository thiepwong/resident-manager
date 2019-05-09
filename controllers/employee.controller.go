package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/thiepwong/resident-manager/models"
	"github.com/thiepwong/resident-manager/services"
	"github.com/thiepwong/smartid/pkg/logger"
)

type EmployeeController struct {
	Ctx     iris.Context
	Service services.EmployeeService
	Result  MvcResult
}

func (c *EmployeeController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("POST", "/register", "PostRegister")
}

func (c *EmployeeController) PostRegister() MvcResult {
	var _signupData = models.Employee{}
	er := c.Ctx.ReadJSON(&_signupData)
	if er != nil {
		logger.LogErr.Println(er)
		return c.Result
	}
	r := c.Service.Register(_signupData.DepartmentId, _signupData.Name, _signupData.Mobile, _signupData.AccountId)
	c.Result.GenerateResult(200, "", r)
	return c.Result
}
