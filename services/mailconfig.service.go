package services

import (
	"time"

	"github.com/thiepwong/resident-manager/common"
	"github.com/thiepwong/resident-manager/models"
	"github.com/thiepwong/resident-manager/repositories"
)

type MailConfigService interface {
	Add(string, string, int64, bool, string) (*models.MailConfig, error)
	GetList(string, int, int, int, int, int, string) (interface{}, error)
	GetById(string) (*models.MailConfig, error)
	// Update(string, string, string, int64, bool, string) (*models.MailConfig, error)
	// Delete(string) (bool, error)
	// SendNotification(m *models.SendNotification) (interface{}, error)
}

type mailconfigServiceImp struct {
	mailconfigRepo repositories.MailConfigRepository
	config         *common.Config
}

func NewMailConfigService(repo repositories.MailConfigRepository, cfg *common.Config) MailConfigService {
	return &mailconfigServiceImp{
		mailconfigRepo: repo,
		config:         cfg,
	}
}

func (s *mailconfigServiceImp) Add(sideId string, title string, publishDate int64, result bool, content string) (*models.MailConfig, error) {
	// _id := uuid.Must(uuid.NewV4())
	if publishDate == 0 {
		publishDate = time.Now().Unix()
	}
	var _noti = &models.MailConfig{}
	return s.mailconfigRepo.Add(_noti)
}

func (s *mailconfigServiceImp) GetList(string, int, int, int, int, int, string) (interface{}, error) {
	return nil, nil
}
func (s *mailconfigServiceImp) GetById(id string) (*models.MailConfig, error) {
	return nil, nil
}
