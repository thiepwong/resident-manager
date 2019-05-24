package repositories

import (
	"errors"

	"github.com/go-pg/pg"
	"github.com/thiepwong/resident-manager/models"
)

type FeedbackRepository interface {
	Add(*models.Feedback) (*models.Feedback, error)
	GetById(string) (*models.FeedbackModel, error)
	GetPagination(string, int, int, string) (*[]models.FeedbackModel, error)
	Update(*models.Feedback) (*models.Feedback, error)
	Delete(*models.Feedback) (bool, error)
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
	r.db.Model(&_feedback).Column("room.*", "Block").Where("room.id=?", id).Select()
	return &_feedback, nil

}

func (r *feedbackRepositoryContext) GetPagination(sideId string, offset int, limit int, orderBy string) (*[]models.FeedbackModel, error) {
	var _feedback []models.FeedbackModel
	if orderBy == "" {
		orderBy = "id DESC"
	}
	r.db.Model(&_feedback).Column("feedback.*", "Side", "Block").Where("feedback.side_id=?", sideId).Order(orderBy).Limit(limit).Offset(offset).Select()
	return &_feedback, nil
}

func (r *feedbackRepositoryContext) Update(m *models.Feedback) (*models.Feedback, error) {

	res, err := r.db.Model(m).Set("status = ?, assigned_employee_id=? , assigned_by =? , due_date=? , actual_finish_date=?", m.Status, m.AssignEmployeeId, m.AssignedBy, m.DueDate, m.ActualFinishDate).Where("id = ?", m.Id).Update()
	if res == nil {
		return nil, errors.New("Feedback id is not found!")
	}
	r.db.Model(m).Where("id=?", m.Id).Select()
	return m, err
}

func (r *feedbackRepositoryContext) Delete(m *models.Feedback) (bool, error) {
	return true, r.db.Delete(m)
}
