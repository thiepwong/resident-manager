package repositories

import (
	"github.com/go-pg/pg"
	"github.com/thiepwong/resident-manager/models"
)

type DepartmentRepository interface {
	Add(*models.Department) (*models.Department, error)
	GetById(string) (*models.DepartmentModel, error)
	GetAll() (*[]models.DepartmentModel, error)
	GetPagination(string, int, int, string) (*[]models.DepartmentModel, error)
	Update(*models.Department) (*models.Department, error)
	Delete(*models.Department) (bool, error)
}

type departmentRepositoryContext struct {
	db *pg.DB
}

func NewDepartmentRepository(db *pg.DB) *departmentRepositoryContext {
	return &departmentRepositoryContext{
		db: db,
	}
}

func (r *departmentRepositoryContext) Add(m *models.Department) (*models.Department, error) {
	return m, r.db.Insert(m)
}

func (r *departmentRepositoryContext) GetById(id string) (*models.DepartmentModel, error) {
	var _dept models.DepartmentModel
	return &_dept, r.db.Model(&_dept).Where("id=?", id).Select()
}

func (r *departmentRepositoryContext) GetAll() (*[]models.DepartmentModel, error) {
	var _dept []models.DepartmentModel

	return &_dept, nil
}

func (r *departmentRepositoryContext) GetPagination(sideId string, offset int, limit int, orderBy string) (*[]models.DepartmentModel, error) {
	var _dept []models.DepartmentModel
	if orderBy == "" {
		orderBy = "id DESC"
	}

	r.db.Model(&_dept).Column("department.*", "Side").Where("side_id=?", sideId).Order(orderBy).Limit(limit).Offset(offset).Select()
	return &_dept, nil
}

func (r *departmentRepositoryContext) Update(m *models.Department) (*models.Department, error) {
	return m, r.db.Update(m)
}

func (r *departmentRepositoryContext) Delete(m *models.Department) (bool, error) {
	return true, r.db.Delete(m)
}
