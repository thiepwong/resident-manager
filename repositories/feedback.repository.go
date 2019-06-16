package repositories

import (
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
	// var _side models.Side
	// r.db.Model(&_side).Where("id=?", _feedback.SideId).Select()
	// _feedback.Block.Side = _side
	// if _feedback.Id == "" {
	// 	return nil, nil
	// }
	return &_feedback, nil

}

func (r *feedbackRepositoryContext) GetPagination(sideId string, offset int, limit int, orderBy string) (*[]models.FeedbackModel, error) {
	var _feedback []models.FeedbackModel
	if orderBy == "" {
		orderBy = "id DESC"
	}
	//	var _side models.Side

	r.db.Model(&_feedback).Column("feedback.*", "Side", "Block", "Room").Where("feedback.side_id=?", sideId).Order(orderBy).Limit(limit).Offset(offset).Select()
	//	r.db.Model(&_side).Where("id=?", sideId).Select()

	// for i := 0; i < len(_feedback); i++ {
	// 	_feedback[i].Block.Side = _side
	// }

	return &_feedback, nil
}

func (r *feedbackRepositoryContext) Update(m *models.Feedback) (*models.Feedback, error) {
	return m, r.db.Update(m)
}

func (r *feedbackRepositoryContext) Delete(m *models.Feedback) (bool, error) {
	return true, r.db.Delete(m)
}
