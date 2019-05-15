package repositories

import (
	"github.com/go-pg/pg"
	"github.com/thiepwong/resident-manager/models"
)

type NotificationRepository interface {
	Add(*models.Notification) (*models.Notification, error)
	GetPagination(int, int, string) (*[]models.Notification, error)
	GetById(id string) (*models.Notification, error)
	Update(*models.Notification) (*models.Notification, error)
	Delete(*models.Notification) (bool, error)
}

type notificationRepositoryContext struct {
	db *pg.DB
}

func NewNotificationRepository(db *pg.DB) *notificationRepositoryContext {
	return &notificationRepositoryContext{
		db: db,
	}
}

func (r *notificationRepositoryContext) Add(m *models.Notification) (*models.Notification, error) {
	return m, r.db.Insert(m)
}

func (r *notificationRepositoryContext) GetPagination(offset int, limit int, orderBy string) (*[]models.Notification, error) {
	var _noti []models.Notification
	if orderBy == "" {
		orderBy = "id DESC"
	}

	r.db.Model(&_noti).Column("send_notification.*").Order(orderBy).Limit(limit).Offset(offset).Select()
	return &_noti, nil
}

func (r *notificationRepositoryContext) GetById(id string) (*models.Notification, error) {
	var _noti models.Notification
	e := r.db.Model(&_noti).Where("id=?", id).Select()
	if e != nil {
		return nil, e
	}
	return &_noti, nil
}

func (r *notificationRepositoryContext) Update(m *models.Notification) (*models.Notification, error) {
	return m, r.db.Update(m)
}

func (r *notificationRepositoryContext) Delete(m *models.Notification) (bool, error) {
	return true, nil
}
