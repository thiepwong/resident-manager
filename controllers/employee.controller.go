package controllers

import (
	"fmt"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris/mvc"
	"github.com/thiepwong/resident-manager/models"
	"github.com/thiepwong/smartid/pkg/logger"
)

// type EmployeeController struct {
// 	Ctx     iris.Context
// 	Service services.EmployeeService
// 	Result  MvcResult
// }

type EmployeeController Controller

func (c *EmployeeController) BeforeActivation(b mvc.BeforeActivation) {
	c.Auth = true
	b.Handle("POST", "/register", "PostRegister")
	b.Handle("GET", "/list", "GetList")
	b.Handle("POST", "/signin", "PostSignin")
}

func (c *EmployeeController) PostRegister() MvcResult {
	var _signupData = models.Employee{}
	_token := c.Ctx.GetHeader("Authorization")
	er := c.Ctx.ReadJSON(&_signupData)
	if er != nil {
		logger.LogErr.Println(er)
		return c.Result
	}

	token, err := jwt.Parse(_token, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return token, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		fmt.Println(err)
	}

	r := c.Service.Register(_signupData.DepartmentId, _signupData.Name, _signupData.Mobile, _signupData.Address, _signupData.AccountId, _signupData.Role, _token)
	c.Result.GenerateResult(200, "", r)
	return c.Result
}

func (c *EmployeeController) GetList() MvcResult {
	_pageIndex, e := strconv.Atoi(c.Ctx.URLParam("page"))
	_pageSize, e := strconv.Atoi(c.Ctx.URLParam("size"))
	_orderBy := c.Ctx.URLParam("order")
	rs, e := c.Service.GetList(_pageIndex, _pageSize, _orderBy)
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
		c.Result.GenerateResult(500, "", e)
		return c.Result
	}
	c.Result.GenerateResult(500, "", rs)
	return c.Result

}
