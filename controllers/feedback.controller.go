package controllers

import (
	"strconv"

	"github.com/thiepwong/resident-manager/models"

	"github.com/kataras/iris/mvc"
	"github.com/thiepwong/resident-manager/services"
)

type FeedbackController struct {
	Controller
	Service services.FeedbackService
}

func (c *FeedbackController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/list/{sideId:string}", "GetList")
	b.Handle("GET", "/detail/{id:string}", "GetById")
	b.Handle("POST", "/update/{id:string}", "PostUpdate")
	b.Handle("GET", "/get-by-emp/{employeeId:string}", "GetByEmp")
	// b.Handle("POST", "/delete/{id:string}", "PostDelete")
}

func (c *FeedbackController) GetList(sideId string) MvcResult {

	if sideId == "" {
		c.Result.GenerateResult(500, "Side id is required for list", nil)
		return c.Result
	}

	_pageIndex, e := strconv.Atoi(c.Ctx.URLParam("page"))
	_pageSize, e := strconv.Atoi(c.Ctx.URLParam("size"))
	_orderBy := c.Ctx.URLParam("order")

	_status, e := strconv.Atoi(c.Ctx.URLParam("status"))
	_residentName := c.Ctx.URLParam("resident")
	_blockId := c.Ctx.URLParam("block")
	_workderName := c.Ctx.URLParam("worker")
	_handlerName := c.Ctx.URLParam("handler")
	_fromDate, e := strconv.Atoi(c.Ctx.URLParam("from"))
	_toDate, e := strconv.Atoi(c.Ctx.URLParam("to"))

	rs, e := c.Service.GetList(sideId, _blockId, _residentName, _workderName, _handlerName, _fromDate, _toDate, _status, _pageIndex, _pageSize, _orderBy)
	if e != nil {
		c.Result.GenerateResult(500, e.Error(), e)
		return c.Result
	}
	c.Result.GenerateResult(200, "", rs)
	return c.Result
}

func (c *FeedbackController) GetById(id string) MvcResult {
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

func (c *FeedbackController) GetByEmp(employeeId string) MvcResult {

	if employeeId == "" {
		c.Result.GenerateResult(500, "Side id is required for list", nil)
		return c.Result
	}

	_pageIndex, e := strconv.Atoi(c.Ctx.URLParam("page"))
	_pageSize, e := strconv.Atoi(c.Ctx.URLParam("size"))
	_orderBy := c.Ctx.URLParam("order")
	rs, e := c.Service.GetListByEmployeeId(employeeId, _pageIndex, _pageSize, _orderBy)
	if e != nil {
		c.Result.GenerateResult(500, e.Error(), e)
		return c.Result
	}
	c.Result.GenerateResult(200, "", rs)
	return c.Result
}

// func (c *FeedbackController) PostAdd() MvcResult {
// 	_feedback := &models.Feedback{}
// 	e := c.Ctx.ReadJSON(_feedback)
// 	if e != nil {
// 		c.Result.GenerateResult(500, e.Error(), e)
// 		return c.Result
// 	}

// 	if _feedback.SideId == "" {
// 		c.Result.GenerateResult(500, "Side Id cannot empty", nil)
// 		return c.Result
// 	}

// 	r, e := c.Service.Add(_feedback.Name, _feedback.SideId)
// 	if e != nil {
// 		c.Result.GenerateResult(500, e.Error(), e)
// 		return c.Result
// 	}

// 	c.Result.GenerateResult(200, "", r)
// 	return c.Result
// }

func (c *FeedbackController) PostUpdate(id string) MvcResult {

	_feedback := &models.Feedback{}

	if id == "" {
		c.Result.GenerateResult(500, "Id is required!", nil)
		return c.Result
	}

	e := c.Ctx.ReadJSON(_feedback)
	if e != nil {
		c.Result.GenerateResult(500, e.Error(), e)
		return c.Result
	}
	r, e := c.Service.Update(id, _feedback.Status, _feedback.AssignEmployeeId, _feedback.AssignedBy, _feedback.DueDate, _feedback.ActualFinishDate)
	if e != nil {
		c.Result.GenerateResult(500, e.Error(), e)
		return c.Result
	}

	c.Result.GenerateResult(200, "", r)
	return c.Result

}

// func (c *FeedbackController) PostDelete(id string) MvcResult {
// 	r, e := c.Service.Delete(id)
// 	if e != nil {
// 		c.Result.GenerateResult(500, e.Error(), e)
// 		return c.Result
// 	}

// 	c.Result.GenerateResult(200, "", r)
// 	return c.Result
// }
