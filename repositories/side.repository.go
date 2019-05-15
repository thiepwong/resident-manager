package repositories

import (
	"github.com/go-pg/pg"
	"github.com/thiepwong/resident-manager/models"
)

type SideRepository interface {
	Add(*models.Side) (*models.Side, error)
	GetById(string) (*models.Side, error)
	GetPagination(int, int, string) (*[]models.Side, error)
	Update(*models.Side) (*models.Side, error)
	Delete(*models.Side) (bool, error)
}

type sideRepositoryContext struct {
	db *pg.DB
}

func NewSideRepository(db *pg.DB) *sideRepositoryContext {
	return &sideRepositoryContext{
		db: db,
	}
}

func (r *sideRepositoryContext) Add(m *models.Side) (*models.Side, error) {
	return m, r.db.Insert(m)
}

func (r *sideRepositoryContext) GetById(id string) (*models.Side, error) {
	var _side models.Side
	return &_side, r.db.Model(&_side).Where("id=?", id).Select()
}

func (r *sideRepositoryContext) GetPagination(offset int, limit int, orderBy string) (*[]models.Side, error) {
	var _side []models.Side
	if orderBy == "" {
		orderBy = "id DESC"
	}
	r.db.Model(&_side).Order(orderBy).Limit(limit).Offset(offset).Select()
	return &_side, nil
}

func (r *sideRepositoryContext) Update(m *models.Side) (*models.Side, error) {
	return m, r.db.Update(m)
}

func (r *sideRepositoryContext) Delete(m *models.Side) (bool, error) {
	return true, r.db.Delete(m)
}
