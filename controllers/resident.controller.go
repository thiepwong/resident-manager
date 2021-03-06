package controllers

import (
	"strconv"

	"github.com/kataras/iris/mvc"
	"github.com/thiepwong/resident-manager/models"
	"github.com/thiepwong/resident-manager/services"
)

type ResidentController struct {
	Controller
	Service services.ResidentService
}

func (c *ResidentController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/list/{sideId:string}", "GetList")
	b.Handle("GET", "/detail/{id:string}", "GetById")
	b.Handle("POST", "/add", "PostAdd")
	b.Handle("POST", "/update", "PostUpdate")
	b.Handle("POST", "/delete/{id:string}", "PostDelete")

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
	_block := c.Ctx.URLParam("block")
	_room := c.Ctx.URLParam("room")
	_name := c.Ctx.URLParam("name")
	_mobile := c.Ctx.URLParam("mobile")
	_email := c.Ctx.URLParam("email")

	rs, e := c.Service.GetList(sideId, _search, _block, _room, _name, _mobile, _email, _pageIndex, _pageSize, _orderBy)
	if e != nil {
		c.Result.GenerateResult(500, e.Error(), e)
		return c.Result
	}
	c.Result.GenerateResult(200, "", rs)
	return c.Result
}

func (c *ResidentController) PostAdd() MvcResult {
	var _resModel models.ResidentModel
	err := c.Ctx.ReadJSON(&_resModel)
	if err != nil {
		c.Result.GenerateResult(500, err.Error(), err)
		return c.Result
	}
	rs, err := c.Service.Add(&_resModel)
	if err != nil {
		c.Result.GenerateResult(500, err.Error(), err)
		return c.Result
	}
	c.Result.GenerateResult(200, "", rs)
	return c.Result
}

func (c *ResidentController) PostUpdate() MvcResult {
	var _resModel models.ResidentModel
	err := c.Ctx.ReadJSON(&_resModel)
	if err != nil {
		c.Result.GenerateResult(500, err.Error(), err)
		return c.Result
	}
	rs, err := c.Service.Update(&_resModel)
	if err != nil {
		c.Result.GenerateResult(500, err.Error(), err)
		return c.Result
	}
	c.Result.GenerateResult(200, "", rs)
	return c.Result
}

func (c *ResidentController) PostDelete(id string) MvcResult {
	if id == "" {
		c.Result.GenerateResult(500, "Id is invalid!", nil)
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
