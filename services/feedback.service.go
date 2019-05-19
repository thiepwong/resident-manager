package services

import (
	"errors"

	uuid "github.com/satori/go.uuid"
	"github.com/thiepwong/resident-manager/models"
	"github.com/thiepwong/resident-manager/repositories"
)

type FeedbackService interface {
	Add(string, string, string) (*models.Feedback, error)
	GetById(string) (*models.FeedbackModel, error)
	GetList(string, int, int, string) (*[]models.FeedbackModel, error)
	Update(string, string, string, string) (*models.Feedback, error)
	Delete(string) (bool, error)
}

type feedbackServiceImp struct {
	feedbackRepo repositories.FeedbackRepository
}

func NewFeedbackService(repo repositories.FeedbackRepository) FeedbackService {
	return &feedbackServiceImp{
		feedbackRepo: repo,
	}
}

func (s *feedbackServiceImp) Add(roomNo string, sideId string, blockId string) (*models.Feedback, error) {

	_id := uuid.Must(uuid.NewV4())
	_feedback := models.Feedback{
		Id: _id.String(),
		// RoomNo:  roomNo,
		// SideId:  sideId,
		// BlockId: blockId,
	}

	return s.feedbackRepo.Add(&_feedback)
}

func (s *feedbackServiceImp) GetById(id string) (*models.FeedbackModel, error) {
	return s.feedbackRepo.GetById(id)
}

func (s *feedbackServiceImp) GetList(sideId string, pageIndex int, pageSize int, orderBy string) (*[]models.FeedbackModel, error) {
	if pageIndex < 1 || pageSize < 1 {
		return nil, errors.New("Page index or page Size is invalid! Please check!")
	}
	var offset int
	offset = (pageIndex - 1) * pageSize
	rs, e := s.feedbackRepo.GetPagination(sideId, offset, pageSize, orderBy)
	if e != nil {
		return nil, e
	}
	return rs, nil
}

func (s *feedbackServiceImp) Update(id string, roomNo string, sideId string, blockId string) (*models.Feedback, error) {
	_feedback := &models.Feedback{
		Id: id,
		// RoomNo:  roomNo,
		// SideId:  sideId,
		// BlockId: blockId,
	}
	return s.feedbackRepo.Update(_feedback)
}

func (s *feedbackServiceImp) Delete(id string) (bool, error) {
	_feedback := &models.Feedback{Id: id}
	return s.feedbackRepo.Delete(_feedback)
}
