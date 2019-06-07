package controllers

import (
	"strconv"

	"github.com/thiepwong/resident-manager/services"

	"github.com/kataras/iris/mvc"
	"github.com/thiepwong/resident-manager/models"
)

type EmployeeController struct {
	Controller
	Service services.EmployeeService
}

func (c *EmployeeController) BeforeActivation(b mvc.BeforeActivation) {
	c.Auth = true
	b.Handle("POST", "/register", "PostRegister")
	b.Handle("GET", "/list/{requestId:string}", "GetList")
	b.Handle("POST", "/signin", "PostSignin")
	b.Handle("POST", "/signup", "PostSignUp")
	b.Handle("POST", "/activate", "PostActivate")
	b.Handle("POST", "/send-otp/{mobile:string}", "PostSendOTP")
	//b.Handle("GET","/detail")
}

func (c *EmployeeController) PostRegister() MvcResult {
	var _signupData = models.Employee{}
	// _token := c.Ctx.GetHeader("Authorization")
	er := c.Ctx.ReadJSON(&_signupData)
	if er != nil {
		return c.Result
	}

	r := c.Service.Register(_signupData.DepartmentId, _signupData.Name, _signupData.Mobile, _signupData.Address, _signupData.AccountId, _signupData.Role, "")
	c.Result.GenerateResult(200, "", r)
	return c.Result
}

func (c *EmployeeController) GetList(requestId string) MvcResult {
	_pageIndex, e := strconv.Atoi(c.Ctx.URLParam("page"))
	_pageSize, e := strconv.Atoi(c.Ctx.URLParam("size"))
	_orderBy := c.Ctx.URLParam("order")
	_isDept, e := strconv.Atoi(c.Ctx.URLParam("type"))
	_role, e := strconv.Atoi(c.Ctx.URLParam("role"))
	var _isDeptmentId bool
	if _isDept == 1 {
		_isDeptmentId = true
	} else {
		_isDeptmentId = false
	}
	rs, e := c.Service.GetList(_isDeptmentId, requestId, _role, _pageIndex, _pageSize, _orderBy)
	if e != nil {
		c.Result.GenerateResult(500, "", e)
		return c.Result
	}

	c.Result.GenerateResult(200, "", rs)
	return c.Result
}

func (c *EmployeeController) PostSignin() MvcResult {
	var _login = models.Signin{}
	e := c.Ctx.ReadJSON(&_login)
	if e != nil {
		return c.Result
	}

	rs, e := c.Service.Signin(_login.Username, _login.Password, _login.System)
	if e != nil {
		c.Result.GenerateResult(500, e.Error(), e)
		return c.Result
	}
	c.Result.GenerateResult(500, "", rs)
	return c.Result

}

func (c *EmployeeController) GetBy(id string) MvcResult {
	if id == "" {
		return c.Result
	}
	rs := c.Service.GetById(id)
	c.Result.GenerateResult(200, "", rs)
	return c.Result
}

func (c *EmployeeController) PostSignUp() MvcResult {
	var signup = models.SignUpModel{}
	er := c.Ctx.ReadJSON(&signup)
	if er != nil {
		c.Result.GenerateResult(500, er.Error(), er)
		return c.Result
	}
	r, er := c.Service.SignUp(&signup)
	if er != nil {
		c.Result.GenerateResult(500, er.Error(), er)
		return c.Result
	}
	c.Result.GenerateResult(200, "", r)
	return c.Result

}

func (c *EmployeeController) PostActivate() MvcResult {
	var activate = models.Activate{}
	er := c.Ctx.ReadJSON(&activate)
	if er != nil {
		c.Result.GenerateResult(500, er.Error(), er)
		return c.Result
	}
	r, er := c.Service.Activate(&activate)
	if er != nil {
		c.Result.GenerateResult(500, er.Error(), er)
		return c.Result
	}
	c.Result.GenerateResult(200, "", r)
	return c.Result

}

func (c *EmployeeController) PostSendOTP(mobile string) MvcResult {

	r, er := c.Service.SendOTP(mobile)
	if er != nil {
		c.Result.GenerateResult(500, er.Error(), er)
		return c.Result
	}
	c.Result.GenerateResult(200, "", r)

	return c.Result

}
