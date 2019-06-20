package controllers

import (
	"strconv"

	"github.com/kataras/iris/mvc"
	"github.com/thiepwong/resident-manager/services"
)

type ResidentController struct {
	Controller
	Service services.ResidentService
}

func (c *ResidentController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/list/{sideId:string}", "GetList")
	b.Handle("GET", "/detail/{id:string}", "GetById")
}

func (c *ResidentController) GetById(id string) MvcResult {
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

func (c *ResidentController) GetList(sideId string) MvcResult {

	if sideId == "" {
		c.Result.GenerateResult(500, "Side id is required for list", nil)
		return c.Result
	}

	_pageIndex, e := strconv.Atoi(c.Ctx.URLParam("page"))
	_pageSize, e := strconv.Atoi(c.Ctx.URLParam("size"))
	_orderBy := c.Ctx.URLParam("order")
	_search := c.Ctx.URLParam("search")
	rs, e := c.Service.GetList(sideId, _search, _pageIndex, _pageSize, _orderBy)
	if e != nil {
		c.Result.GenerateResult(500, e.Error(), e)
		return c.Result
	}
	c.Result.GenerateResult(200, "", rs)
	return c.Result
}
