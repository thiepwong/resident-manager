package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/thiepwong/resident-manager/common"
	"github.com/thiepwong/resident-manager/models"
	"github.com/thiepwong/resident-manager/repositories"
)

type NotificationService interface {
	Add(string, string, int64, bool, string) (*models.Notification, error)
	GetList(string, int, int, string) (*[]models.NotificationModel, error)
	GetById(string) (*models.NotificationModel, error)
	Update(string, string, string, int64, bool, string) (*models.Notification, error)
	Delete(string) (bool, error)
	SendNotification(m *models.SendNotification) (interface{}, error)
}

type notificationServiceImp struct {
	notiRepo repositories.NotificationRepository
	config   *common.Config
}

func NewNotificationService(repo repositories.NotificationRepository, cfg *common.Config) NotificationService {
	return &notificationServiceImp{
		notiRepo: repo,
		config:   cfg,
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

func (s *notificationServiceImp) GetList(sideId string, pageIndex int, pageSize int, orderBy string) (*[]models.NotificationModel, error) {
	if pageIndex < 1 || pageSize < 1 {
		return nil, errors.New("Page index or page Size is invalid! Please check!")
	}
	if sideId == "" {
		return nil, errors.New("Side id is required for notification list!")
	}
	var offset int
	offset = (pageIndex - 1) * pageSize
	rs, e := s.notiRepo.GetPagination(sideId, offset, pageSize, orderBy)
	if e != nil {
		return nil, e
	}
	return rs, nil
}

func (s *notificationServiceImp) GetById(id string) (*models.NotificationModel, error) {

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
	m := models.Notification{Id: id}
	return s.notiRepo.Delete(&m)
}

func (s *notificationServiceImp) SendNotification(m *models.SendNotification) (interface{}, error) {
	message := map[string]interface{}{
		"to":           "/topics/" + m.Topic,
		"collapse_key": "type_a",
		"notification": map[string]interface{}{
			"body":  m.Content,
			"title": m.Title},
		"data": map[string]interface{}{
			"body":    m.Content,
			"title":   m.Title,
			"channel": m.Title,
			"key_2":   "Value for key_2"}}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	req, e := http.NewRequest("POST", s.config.Option.FireBaseUrl, bytes.NewBuffer(bytesRepresentation))
	if e != nil {
		log.Fatal("Loi roi")
		return nil, err
	}

	req.Header.Set("Content-Type", "Application/json")
	req.Header.Set("Authorization", s.config.Option.FireBaseToken)

	// Do the request
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	var result map[string]interface{}
	json.NewDecoder(response.Body).Decode(&result)
	if result["error"] != nil {
		s := fmt.Sprintf("%s", result["error"])
		return nil, errors.New(s)
	}
	return result, nil
}
