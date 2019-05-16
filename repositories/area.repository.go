package repositories

import (
	"github.com/go-pg/pg"
	"github.com/thiepwong/resident-manager/models"
)

type AreaRepository interface {
	Add(*models.Area) (*models.Area, error)
	GetById(string) (*models.AreaModel, error)
	GetPagination(string, int, int, string) (*[]models.AreaModel, error)
	Update(*models.Area) (*models.Area, error)
	Delete(*models.Area) (bool, error)
}

type areaRepositoryContext struct {
	db *pg.DB
}

func NewAreaRepository(db *pg.DB) *areaRepositoryContext {
	return &areaRepositoryContext{
		db: db,
	}
}

func (r *areaRepositoryContext) Add(m *models.Area) (*models.Area, error) {
	return m, r.db.Insert(m)
}

func (r *areaRepositoryContext) GetById(id string) (*models.AreaModel, error) {
	var _area models.AreaModel
	return &_area, r.db.Model(&_area).Column("area.*", "Side").Where("area.id=?", id).Select()
}

func (r *areaRepositoryContext) GetPagination(sideId string, offset int, limit int, orderBy string) (*[]models.AreaModel, error) {
	var _area []models.AreaModel
	if orderBy == "" {
		orderBy = "id DESC"
	}
	r.db.Model(&_area).Column("area.*", "Side").Where("side_id=?", sideId).Order(orderBy).Limit(limit).Offset(offset).Select()
	return &_area, nil
}

func (r *areaRepositoryContext) Update(m *models.Area) (*models.Area, error) {
	return m, r.db.Update(m)
}

func (r *areaRepositoryContext) Delete(m *models.Area) (bool, error) {
	return true, r.db.Delete(m)
}
