package repositories

import (
	"github.com/go-pg/pg"
	"github.com/thiepwong/resident-manager/models"
)

type DepartmentRepository interface {
	Add(*models.Department) (*models.Department, error)
	GetById(string) (*models.Department, error)
	GetAll() (*[]models.Department, error)
	GetPagination(int, int, string) (*[]models.Department, error)
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

func (r *departmentRepositoryContext) GetById(id string) (*models.Department, error) {
	var _dept models.Department
	return &_dept, r.db.Model(&_dept).Where("id=?", id).Select()
}

func (r *departmentRepositoryContext) GetAll() (*[]models.Department, error) {
	var _dept []models.Department

	return &_dept, nil
}

func (r *departmentRepositoryContext) GetPagination(offset int, limit int, orderBy string) (*[]models.Department, error) {
	var _dept []models.Department
	if orderBy == "" {
		orderBy = "id DESC"
	}

	r.db.Model(&_dept).Column("department.*").Order(orderBy).Limit(limit).Offset(offset).Select()
	return &_dept, nil
}

func (r *departmentRepositoryContext) Update(m *models.Department) (*models.Department, error) {
	return m, r.db.Update(m)
}

func (r *departmentRepositoryContext) Delete(m *models.Department) (bool, error) {
	return true, r.db.Delete(m)
}
