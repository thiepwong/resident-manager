package controllers

import (
	"strconv"

	"github.com/thiepwong/resident-manager/models"

	"github.com/kataras/iris/mvc"
	"github.com/thiepwong/resident-manager/services"
)

type NotificationController struct {
	Controller
	Service services.NotificationService
}

func (c *NotificationController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/list/{sideId}", "GetList")
	b.Handle("GET", "/detail/{id:string}", "GetById")
	b.Handle("POST", "/add", "PostAdd")
	b.Handle("POST", "/update/{id:string}", "PostUpdate")
}

func (c *NotificationController) GetList(sideId string) MvcResult {
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

func (c *NotificationController) GetById(id string) MvcResult {
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

func (c *NotificationController) PostAdd() MvcResult {
	_noti := &models.Notification{}
	e := c.Ctx.ReadJSON(_noti)
	if e != nil {
		c.Result.GenerateResult(500, e.Error(), e)
		return c.Result
	}
	r, e := c.Service.Add(_noti.SideId, _noti.Title, _noti.PublishDate, _noti.SendResult, _noti.Content)
	if e != nil {
		c.Result.GenerateResult(500, e.Error(), e)
		return c.Result
	}

	c.Result.GenerateResult(200, "", r)
	return c.Result
}

func (c *NotificationController) PostUpdate(id string) MvcResult {
	_noti := &models.Notification{}
	e := c.Ctx.ReadJSON(_noti)
	if e != nil {
		c.Result.GenerateResult(500, e.Error(), e)
		return c.Result
	}
	r, e := c.Service.Update(id, _noti.SideId, _noti.Title, _noti.PublishDate, _noti.SendResult, _noti.Content)
	if e != nil {
		c.Result.GenerateResult(500, e.Error(), e)
		return c.Result
	}

	c.Result.GenerateResult(200, "", r)
	return c.Result
}
