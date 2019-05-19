package repositories

import (
	"github.com/go-pg/pg"
	"github.com/thiepwong/resident-manager/models"
)

type NotificationRepository interface {
	Add(*models.Notification) (*models.Notification, error)
	GetPagination(string, int, int, string) (*[]models.NotificationModel, error)
	GetById(id string) (*models.NotificationModel, error)
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

func (r *notificationRepositoryContext) GetPagination(sideId string, offset int, limit int, orderBy string) (*[]models.NotificationModel, error) {
	var _noti []models.NotificationModel
	if orderBy == "" {
		orderBy = "title DESC"
	}

	r.db.Model(&_noti).Column("send_notification.*", "Side").Where("side_id=?", sideId).Order(orderBy).Limit(limit).Offset(offset).Select()
	return &_noti, nil
}

func (r *notificationRepositoryContext) GetById(id string) (*models.NotificationModel, error) {
	var _noti models.NotificationModel
	e := r.db.Model(&_noti).Column("send_notification.*", "Side").Where("send_notification.id=?", id).Select()
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
