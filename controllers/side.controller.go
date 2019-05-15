package controllers

import (
	"strconv"

	"github.com/thiepwong/resident-manager/models"

	"github.com/kataras/iris/mvc"
	"github.com/thiepwong/resident-manager/services"
)

type SideController struct {
	Controller
	Service services.SideService
}

func (c *SideController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/list", "GetList")
	b.Handle("GET", "/detail/{id:string}", "GetById")
	b.Handle("POST", "/add", "PostAdd")
	b.Handle("POST", "/update/{id:string}", "PostUpdate")
	b.Handle("POST", "/delete/{id:string}", "PostDelete")
}

func (c *SideController) GetList() MvcResult {
	_pageIndex, e := strconv.Atoi(c.Ctx.URLParam("page"))
	_pageSize, e := strconv.Atoi(c.Ctx.URLParam("size"))
	_orderBy := c.Ctx.URLParam("order")
	rs, e := c.Service.GetList(_pageIndex, _pageSize, _orderBy)
	if e != nil {
		c.Result.GenerateResult(500, e.Error(), e)
		return c.Result
	}
	c.Result.GenerateResult(200, "", rs)
	return c.Result
}

func (c *SideController) GetById(id string) MvcResult {
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

func (c *SideController) PostAdd() MvcResult {
	_side := &models.Side{}
	e := c.Ctx.ReadJSON(_side)
	if e != nil {
		c.Result.GenerateResult(500, e.Error(), e)
		return c.Result
	}
	r, e := c.Service.Add(_side.Name, _side.Address, _side.Ip, _side.Cover, _side.Hotline)
	if e != nil {
		c.Result.GenerateResult(500, e.Error(), e)
		return c.Result
	}

	c.Result.GenerateResult(200, "", r)
	return c.Result
}

func (c *SideController) PostUpdate(id string) MvcResult {

	_side := &models.Side{}
	e := c.Ctx.ReadJSON(_side)
	if e != nil {
		c.Result.GenerateResult(500, e.Error(), e)
		return c.Result
	}
	r, e := c.Service.Update(id, _side.Name, _side.Address, _side.Ip, _side.Cover, _side.Hotline)
	if e != nil {
		c.Result.GenerateResult(500, e.Error(), e)
		return c.Result
	}

	c.Result.GenerateResult(200, "", r)
	return c.Result

}

func (c *SideController) PostDelete(id string) MvcResult {
	r, e := c.Service.Delete(id)
	if e != nil {
		c.Result.GenerateResult(500, e.Error(), e)
		return c.Result
	}

	c.Result.GenerateResult(200, "", r)
	return c.Result
}
