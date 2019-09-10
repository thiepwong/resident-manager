package repositories

import (
	"github.com/go-pg/pg"
	"github.com/thiepwong/resident-manager/models"
)

type MailConfigRepository interface {
	Add(*models.MailConfig) (*models.MailConfig, error)
	// GetPagination(string, int, int, int, int, int, string) (interface{}, error)
	GetById(id string) (*models.MailConfig, error)
	Update(*models.MailConfig) (*models.MailConfig, error)
	Delete(*models.MailConfig) (bool, error)
}

type mailconfigRepositoryContext struct {
	db *pg.DB
}

func NewMailConfigRepository(db *pg.DB) *mailconfigRepositoryContext {
	return &mailconfigRepositoryContext{
		db: db,
	}
}

func (r *mailconfigRepositoryContext) Add(m *models.MailConfig) (*models.MailConfig, error) {
	return m, r.db.Insert(m)
}

func (r *mailconfigRepositoryContext) GetById(id string) (*models.MailConfig, error) {
	var _noti models.MailConfig
	e := r.db.Model(&_noti).Column("mail_config.*", "Side").Where("mail_config.id=?", id).Select()
	if e != nil {
		return nil, e
	}
	return &_noti, nil
}
