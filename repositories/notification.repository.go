package repositories

import (
	"github.com/go-pg/pg"
	"github.com/thiepwong/resident-manager/models"
)

type NotificationRepository interface {
	Add(*models.Notification) (*models.Notification, error)
	GetPagination(int, int, string) (*[]models.Notification, error)
	Detail(*models.Notification) (*models.Notification, error)
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
	return m, nil
}

func (r *notificationRepositoryContext) GetList(pageIndex int, pageSize int, orderBy string) (*models.Notification, error) {
	var _noti models.Notification

	return &_noti, nil
}

func (r *notificationRepositoryContext) Detail(m *models.Notification) (*models.Notification, error) {
	var _noti models.Notification
	return &_noti, nil
}

func (r *notificationRepositoryContext) Update(m *models.Notification) (*models.Notification, error) {
	return m, nil
}

func (r *notificationRepositoryContext) Delete(m *models.Notification) (bool, error) {
	return true, nil
}
