package services

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/thiepwong/resident-manager/models"
	"github.com/thiepwong/resident-manager/repositories"
)

type NotificationService interface {
	Add(string, string, int64, bool, string) (*models.Notification, error)
	GetList(int, int, string) (*[]models.Notification, error)
	GetById(string) (*models.Notification, error)
	Update(string, string, string, int64, bool, string) (*models.Notification, error)
	Delete(string) (bool, error)
}

type notificationServiceImp struct {
	notiRepo repositories.NotificationRepository
}

func NewNotificationService(repo repositories.NotificationRepository) NotificationService {
	return &notificationServiceImp{
		notiRepo: repo,
	}
}

func (s *notificationServiceImp) Add(sideId string, title string, publishDate int64, result bool, content string) (*models.Notification, error) {
	_id := uuid.Must(uuid.NewV4())
	if publishDate == 0 {
		publishDate = time.Now().Unix()
	}
	var _noti = &models.Notification{
		Id:          _id.String(),
		SideId:      sideId,
		Title:       title,
		PublishDate: publishDate,
		SendResult:  result,
		Content:     content,
	}
	return s.notiRepo.Add(_noti)
}

func (s *notificationServiceImp) GetList(pageIndex int, pageSize int, orderBy string) (*[]models.Notification, error) {
	if pageIndex < 1 || pageSize < 1 {
		return nil, errors.New("Page index or page Size is invalid! Please check!")
	}
	var offset int
	offset = (pageIndex - 1) * pageSize
	rs, e := s.notiRepo.GetPagination(offset, pageSize, orderBy)
	if e != nil {
		return nil, e
	}
	return rs, nil
}

func (s *notificationServiceImp) GetById(id string) (*models.Notification, error) {

	rs, e := s.notiRepo.GetById(id)
	if e != nil {
		return nil, e
	}
	return rs, nil
}

func (s *notificationServiceImp) Update(id string, sideId string, title string, publishDate int64, sendResult bool, content string) (*models.Notification, error) {

	var _noti = models.Notification{
		Id:          id,
		SideId:      sideId,
		Title:       title,
		PublishDate: publishDate,
		SendResult:  sendResult,
		Content:     content,
	}

	return s.notiRepo.Update(&_noti)

}

func (s *notificationServiceImp) Delete(id string) (bool, error) {
	return true, nil
}
