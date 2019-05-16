package repositories

import (
	"github.com/go-pg/pg"
	"github.com/thiepwong/resident-manager/models"
)

type ContactRepository interface {
	Add(*models.Contact) (*models.Contact, error)
	GetById(string) (*models.ContactModel, error)
	GetPagination(string, int, int, string) (*[]models.ContactModel, error)
	Update(*models.Contact) (*models.Contact, error)
	Delete(*models.Contact) (bool, error)
}

type contactRepositoryContext struct {
	db *pg.DB
}

func NewContactRepository(db *pg.DB) *contactRepositoryContext {
	return &contactRepositoryContext{
		db: db,
	}
}

func (r *contactRepositoryContext) Add(m *models.Contact) (*models.Contact, error) {
	return m, r.db.Insert(m)
}

func (r *contactRepositoryContext) GetById(id string) (*models.ContactModel, error) {
	var _Contact models.ContactModel
	return &_Contact, r.db.Model(&_Contact).Column("contact.*", "Department").Where("contact.id=?", id).Select()
}

func (r *contactRepositoryContext) GetPagination(department_id string, offset int, limit int, orderBy string) (*[]models.ContactModel, error) {
	var _Contact []models.ContactModel
	if orderBy == "" {
		orderBy = "id DESC"
	}
	r.db.Model(&_Contact).Column("contact.*", "Department").Where("department_id=?", department_id).Order(orderBy).Limit(limit).Offset(offset).Select()
	return &_Contact, nil
}

func (r *contactRepositoryContext) Update(m *models.Contact) (*models.Contact, error) {
	return m, r.db.Update(m)
}

func (r *contactRepositoryContext) Delete(m *models.Contact) (bool, error) {
	return true, r.db.Delete(m)
}
