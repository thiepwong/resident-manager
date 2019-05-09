package repositories

import (
	"fmt"

	"github.com/go-pg/pg"
	"github.com/thiepwong/resident-manager/models"
)

type EmployeeRepository interface {
	Register(*models.Employee) (*models.Employee, error)
	GetById(string) (*models.Employee, error)
	GetAll() (*[]models.Employee, error)
	Update(*models.Employee) (*models.Employee, error)
}

type employeeRepositoryContext struct {
	db *pg.DB
}

func NewEmployeeRepository(db *pg.DB) *employeeRepositoryContext {
	return &employeeRepositoryContext{
		db: db,
	}
}

func (emp *employeeRepositoryContext) Register(employee *models.Employee) (*models.Employee, error) {

	e := emp.db.Insert(employee)
	fmt.Print(e.Error())
	return employee, nil
}

func (emp *employeeRepositoryContext) GetById(id string) (*models.Employee, error) {
	var _emp = &models.Employee{}
	return _emp, nil
}

func (emp *employeeRepositoryContext) GetAll() (*[]models.Employee, error) {

	return nil, nil
}

func (emp *employeeRepositoryContext) Update(employee *models.Employee) (*models.Employee, error) {
	return employee, nil
}
