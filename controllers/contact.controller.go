package controllers

import (
	"strconv"

	"github.com/kataras/iris/mvc"
	"github.com/thiepwong/resident-manager/models"
	"github.com/thiepwong/resident-manager/services"
)

type ContactController struct {
	Controller
	Service services.ContactService
}

func (c *ContactController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/list/{departmentId:string}", "GetList")
	b.Handle("GET", "/detail/{id:string}", "GetById")
	b.Handle("POST", "/add", "PostAdd")
	b.Handle("POST", "/update/{id:string}", "PostUpdate")
	b.Handle("POST", "/delete/{id:string}", "PostDelete")
}

func (c *ContactController) GetList(departmentId string) MvcResult {

	if departmentId == "" {
		c.Result.GenerateResult(500, "Side id is required for list", nil)
		return c.Result
	}

	_pageIndex, e := strconv.Atoi(c.Ctx.URLParam("page"))
	_pageSize, e := strconv.Atoi(c.Ctx.URLParam("size"))
	_orderBy := c.Ctx.URLParam("order")
	rs, e := c.Service.GetList(departmentId, _pageIndex, _pageSize, _orderBy)
	if e != nil {
		c.Result.GenerateResult(500, e.Error(), e)
		return c.Result
	}
	c.Result.GenerateResult(200, "", rs)
	return c.Result
}

func (c *ContactController) GetById(id string) MvcResult {
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

func (c *ContactController) PostAdd() MvcResult {
	_contact := &models.Contact{}
	e := c.Ctx.ReadJSON(_contact)
	if e != nil {
		c.Result.GenerateResult(500, e.Error(), e)
		return c.Result
	}

	if _contact.DepartmentId == "" {
		c.Result.GenerateResult(500, "Side Id cannot empty", nil)
		return c.Result
	}

	r, e := c.Service.Add(_contact.Name, _contact.DepartmentId)
	if e != nil {
		c.Result.GenerateResult(500, e.Error(), e)
		return c.Result
	}

	c.Result.GenerateResult(200, "", r)
	return c.Result
}

func (c *ContactController) PostUpdate(id string) MvcResult {

	_contact := &models.Contact{}
	e := c.Ctx.ReadJSON(_contact)
	if e != nil {
		c.Result.GenerateResult(500, e.Error(), e)
		return c.Result
	}
	r, e := c.Service.Update(id, _contact.Name, _contact.DepartmentId, _contact.PhoneNumber)
	if e != nil {
		c.Result.GenerateResult(500, e.Error(), e)
		return c.Result
	}

	c.Result.GenerateResult(200, "", r)
	return c.Result

}

func (c *ContactController) PostDelete(id string) MvcResult {
	r, e := c.Service.Delete(id)
	if e != nil {
		c.Result.GenerateResult(500, e.Error(), e)
		return c.Result
	}

	c.Result.GenerateResult(200, "", r)
	return c.Result
}
