package repositories

import (
	"fmt"

	"github.com/go-pg/pg"
	"github.com/thiepwong/resident-manager/models"
)

type NotificationRepository interface {
	Add(*models.Notification) (*models.Notification, error)
	GetPagination(string, int, int, bool, int, int, string) (interface{}, error)
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

func (r *notificationRepositoryContext) GetPagination(sideId string, fromDate int, toDate int, status bool, offset int, limit int, orderBy string) (interface{}, error) {
	var _noti []models.NotificationQueryModel
	var _result models.ModelResult
	var count int

	if orderBy == "" {
		orderBy = "title DESC"
	}

	count_query := fmt.Sprintf("select count(sn.id)  from resident.send_notification sn left join resident.side rs on rs.id = sn.side_id where sn.side_id = '%s'	and (sn.publish_date >= %d or %d = 0) and (sn.publish_date <= %d or %d=0 ) and (sn.send_success = %t or %t = null) ", sideId, fromDate, fromDate, toDate, toDate, status, status)
	select_query := fmt.Sprintf("select sn.id, sn.title, sn.publish_date, sn.send_success, sn.body, rs.id side_id, rs.NAME side_name, rs.hotline,rs.address,rs.cover_photos from resident.send_notification sn left join resident.side rs on rs.id = sn.side_id where sn.side_id = '%s'	and (sn.publish_date >= %d or %d = 0) and (sn.publish_date <= %d or %d=0 ) and (sn.send_success = %t or %t = null)  order by  %s offset %d limit %d", sideId, fromDate, fromDate, toDate, toDate, status, status, orderBy, offset, limit)
	_, e := r.db.Query(&count, count_query)
	_, e = r.db.Query(&_noti, select_query)

	if e != nil {
		return nil, e
	}
	_result.TotalRecord = count
	_result.Rows = &_noti
	return _result, nil

	// r.db.Model(&_noti).Column("send_notification.*", "Side").Where("side_id=?", sideId).Order(orderBy).Limit(limit).Offset(offset).Select()
	// return &_noti, nil
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
	res, e := r.db.Model(m).Where("id = ?", m.Id).Delete()
	if e != nil {
		return false, e
	}
	if res != nil {

	}
	return true, nil
}
