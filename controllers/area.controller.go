package controllers

import (
	"strconv"

	"github.com/kataras/iris/mvc"
	"github.com/thiepwong/resident-manager/models"
	"github.com/thiepwong/resident-manager/services"
)

type AreaController struct {
	Controller
	Service services.AreaService
}

func (c *AreaController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/list/{sideId:string}", "GetList")
	b.Handle("GET", "/detail/{id:string}", "GetById")
	b.Handle("POST", "/add", "PostAdd")
	b.Handle("POST", "/update/{id:string}", "PostUpdate")
	b.Handle("POST", "/delete/{id:string}", "PostDelete")
}

func (c *AreaController) GetList(sideId string) MvcResult {

	if sideId == "" {
		c.Result.GenerateResult(500, "Side id is required for list", nil)
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

func (c *AreaController) GetById(id string) MvcResult {
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

func (c *AreaController) PostAdd() MvcResult {
	_area := &models.Area{}
	e := c.Ctx.ReadJSON(_area)
	if e != nil {
		c.Result.GenerateResult(500, e.Error(), e)
		return c.Result
	}

	if _area.SideId == "" {
		c.Result.GenerateResult(500, "Side Id cannot empty", nil)
		return c.Result
	}

	r, e := c.Service.Add(_area.Name, _area.SideId)
	if e != nil {
		c.Result.GenerateResult(500, e.Error(), e)
		return c.Result
	}

	c.Result.GenerateResult(200, "", r)
	return c.Result
}

func (c *AreaController) PostUpdate(id string) MvcResult {

	_area := &models.Area{}
	e := c.Ctx.ReadJSON(_area)
	if e != nil {
		c.Result.GenerateResult(500, e.Error(), e)
		return c.Result
	}
	r, e := c.Service.Update(id, _area.Name, _area.SideId)
	if e != nil {
		c.Result.GenerateResult(500, e.Error(), e)
		return c.Result
	}

	c.Result.GenerateResult(200, "", r)
	return c.Result

}

func (c *AreaController) PostDelete(id string) MvcResult {
	r, e := c.Service.Delete(id)
	if e != nil {
		c.Result.GenerateResult(500, e.Error(), e)
		return c.Result
	}

	c.Result.GenerateResult(200, "", r)
	return c.Result
}
