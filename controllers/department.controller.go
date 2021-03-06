package controllers

import (
	"strconv"

	"github.com/kataras/iris/mvc"
	"github.com/thiepwong/resident-manager/models"
	"github.com/thiepwong/resident-manager/services"
)

type DepartmentController struct {
	Controller
	Service services.DepartmentService
}

func (c *DepartmentController) BeforeActivation(b mvc.BeforeActivation) {
	c.Auth = true

	b.Handle("POST", "/add", "PostAdd")
	b.Handle("GET", "/list/{sideId:string}", "GetList")
	b.Handle("POST", "/detail/{id:string}", "GetById")
	b.Handle("POST", "/update/{id:string}", "PostUpdate")
	b.Handle("POST", "/delete/{id:string}", "PostDelete")
}

func (c *DepartmentController) PostAdd() MvcResult {
	var _dept models.Department
	err := c.Ctx.ReadJSON(&_dept)
	if err != nil {
		c.Result.GenerateResult(500, err.Error(), err)
		return c.Result
	}
	rs, err := c.Service.Add(&_dept)
	if err != nil {
		c.Result.GenerateResult(500, err.Error(), err)
		return c.Result
	}
	c.Result.GenerateResult(200, "", rs)
	return c.Result
}

func (c *DepartmentController) GetList(sideId string) MvcResult {
	if sideId == "" {
		c.Result.GenerateResult(500, "Side id is required!", nil)
		return c.Result
	}

	_pageIndex, e := strconv.Atoi(c.Ctx.URLParam("page"))
	_pageSize, e := strconv.Atoi(c.Ctx.URLParam("size"))
	_orderBy := c.Ctx.URLParam("order")
	rs, e := c.Service.GetList(sideId, _pageIndex, _pageSize, _orderBy)
	if e != nil {
		c.Result.GenerateResult(500, e.Error(), e)
		return c.Result
	}

	c.Result.GenerateResult(200, "", rs)
	return c.Result
}

func (c *DepartmentController) GetById(id string) MvcResult {
	if id == "" {
		return c.Result
	}
	rs, e := c.Service.GetById(id)
	if e != nil {
		c.Result.GenerateResult(500, e.Error(), e)
		return c.Result
	}
	c.Result.GenerateResult(200, "", rs)
	return c.Result
}

func (c *DepartmentController) PostUpdate(id string) MvcResult {
	var _dept models.Department
	err := c.Ctx.ReadJSON(&_dept)
	if err != nil {
		c.Result.GenerateResult(500, err.Error(), err)
		return c.Result
	}
	rs, err := c.Service.Update(id, _dept.Name, _dept.SideId)
	if err != nil {
		c.Result.GenerateResult(500, err.Error(), err)
		return c.Result
	}
	c.Result.GenerateResult(200, "", rs)
	return c.Result
}

func (c *DepartmentController) PostDelete(id string) MvcResult {
	var _dept models.Department
	err := c.Ctx.ReadJSON(&_dept)
	if err != nil {
		c.Result.GenerateResult(500, err.Error(), err)
		return c.Result
	}
	rs, err := c.Service.Delete(id)
	if err != nil {
		c.Result.GenerateResult(500, err.Error(), err)
		return c.Result
	}
	c.Result.GenerateResult(200, "", rs)
	return c.Result
}
