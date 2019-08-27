package repositories

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-pg/pg"
	"github.com/thiepwong/resident-manager/models"
)

type FeedbackRepository interface {
	Add(*models.Feedback) (*models.Feedback, error)
	GetById(string) (*models.FeedbackModel, error)
	GetPagination(string, string, string, string, string, int, int, int, int, int, string) (interface{}, error)
	Update(*models.Feedback) (*models.Feedback, error)
	Delete(*models.Feedback) (bool, error)
	GetListByEmployeeId(string, int, int, string) (interface{}, error)
}

type feedbackRepositoryContext struct {
	db *pg.DB
}

func NewFeedbackRepository(db *pg.DB) *feedbackRepositoryContext {
	return &feedbackRepositoryContext{
		db: db,
	}
}

func (r *feedbackRepositoryContext) Add(m *models.Feedback) (*models.Feedback, error) {
	return m, r.db.Insert(m)
}

func (r *feedbackRepositoryContext) GetById(id string) (*models.FeedbackModel, error) {

	var _feedback models.FeedbackModel
	r.db.Model(&_feedback).Column("feedback.*", "Side", "Block", "Room", "Employee").Where("feedback.id=?", id).Select()
	return &_feedback, nil

}

func (r *feedbackRepositoryContext) GetPagination(sideId string, blockId string, residentName string, workerName string, handlerName string, fromDate int, toDate int, status int, offset int, limit int, orderBy string) (interface{}, error) {
	var _feedback []models.FeedbackQueryModel
	var _result models.ModelResult
	var count int
	if orderBy == "" {
		orderBy = "id DESC"
	}
	// count, e := r.db.Model(&_feedback).Column("feedback.*", "Resident", "Side", "Block", "Room", "Employee").Where("feedback.side_id=?", sideId).Where("feedback.block_id = ? or ? = '' ", blockId, blockId).Where("feedback.employee_id = ? or ? = ''", employeeId, employeeId).Where("feedback.status =? or ? = -1", status, status).Order(orderBy).Count()
	// r.db.Model(&_feedback).Column("feedback.*", "Resident", "Side", "Block", "Room", "Employee").Where("feedback.side_id=?", sideId).Where("feedback.block_id = ? or ? =  '' ", blockId, blockId).Where("feedback.employee_id = ? or ? = ''", employeeId, employeeId).Where("feedback.status =? or ? = -1", status, status).Order(orderBy).Limit(limit).Offset(offset).Select()
	// if e != nil {
	// 	return nil, e
	// }residentNamego

	count_query := fmt.Sprintf("select count(fb.id) from resident.feedback fb left join resident.resident re on re.id = fb.resident_id left join resident.room ro on ro.id = fb.room_id left join resident.employee em on em.id = fb.employee_id left join resident.block bl on bl.id = fb.block_id left join resident.employee emp on emp.id = fb.assigned_by  where fb.side_id='%s' and (lower(fb.block_id) = '%s' or '%s'='') and (lower(re.full_name) like '%s' or '%s' ='') and (lower(em.NAME) like '%s' or '%s' = '') and (lower(emp.NAME) like '%s' or '%s' ='') and (fb.created >= %d or %d=0) and (fb.created <= %d or %d=0) and (fb.status = %d or %d=-5)", sideId, blockId, blockId, "%"+strings.ToLower(residentName)+"%", strings.ToLower(residentName), "%"+strings.ToLower(workerName)+"%", strings.ToLower(workerName), "%"+strings.ToLower(handlerName)+"%", strings.ToLower(handlerName), fromDate, fromDate, toDate, toDate, status, status)
	select_query := fmt.Sprintf("select fb.id, fb.title,fb.content,fb.images, fb.status,fb.due_date,fb.actual_finish_date,fb.position_note,fb.created,re.full_name,re.phone_no, ro.room_no,em.id worker_id, em.NAME worker_name,em.mobile,bl.id block_id, bl.NAME block_name, fb.side_id, emp.NAME handler_name from resident.feedback fb left join resident.resident re on re.id = fb.resident_id left join resident.room ro on ro.id = fb.room_id left join resident.employee em on em.id = fb.employee_id left join resident.block bl on bl.id = fb.block_id left join resident.employee emp on emp.id = fb.assigned_by where fb.side_id='%s' and (lower(fb.block_id) = '%s' or '%s'='') and (lower(re.full_name) like '%s' or '%s' ='')	and (lower(em.NAME) like '%s' or '%s' = '')	and (lower(emp.NAME) like '%s' or '%s' ='')	and (fb.created >= %d or %d=0)	and (fb.created <= %d or %d=0) and (fb.status = %d or %d=-5) order by  %s offset %d limit %d", sideId, blockId, blockId, "%"+strings.ToLower(residentName)+"%", strings.ToLower(residentName), "%"+strings.ToLower(workerName)+"%", strings.ToLower(workerName), "%"+strings.ToLower(handlerName)+"%", strings.ToLower(handlerName), fromDate, fromDate, toDate, toDate, status, status, orderBy, offset, limit)
	_, e := r.db.Query(&count, count_query)
	_, e = r.db.Query(&_feedback, select_query)

	if e != nil {
		return nil, e
	}

	_result.TotalRecord = count
	_result.Rows = &_feedback
	return _result, nil
}

func (r *feedbackRepositoryContext) Update(m *models.Feedback) (*models.Feedback, error) {

	res, err := r.db.Model(m).Set("status = ?, employee_id=? , assigned_by =? , due_date=? , actual_finish_date=?", m.Status, m.AssignEmployeeId, m.AssignedBy, m.DueDate, m.ActualFinishDate).Where("id = ?", m.Id).Update()
	if res == nil {
		return nil, errors.New("Feedback id is not found!")
	}
	r.db.Model(m).Where("id=?", m.Id).Select()
	return m, err
}

func (r *feedbackRepositoryContext) Delete(m *models.Feedback) (bool, error) {
	return true, r.db.Delete(m)
}

func (r *feedbackRepositoryContext) GetListByEmployeeId(employeeId string, offset int, limit int, orderBy string) (interface{}, error) {
	if orderBy == "" {
		orderBy = "feedback.created DESC"
	}
	var _feedback []models.FeedbackModel
	_count, e := r.db.Model(&_feedback).Column("feedback.*", "Resident", "Side", "Block", "Room", "Employee").Where("feedback.employee_id=?", employeeId).Order(orderBy).Count()
	r.db.Model(&_feedback).Column("feedback.*", "Resident", "Side", "Block", "Room", "Employee").Where("feedback.employee_id=?", employeeId).Order(orderBy).Limit(limit).Offset(offset).Select()

	if e != nil {
		return nil, e
	}
	result := &models.ModelResult{TotalRecord: _count, Rows: &_feedback}

	return result, nil
}
