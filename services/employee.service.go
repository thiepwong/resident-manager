package services

import (
	"github.com/thiepwong/resident-manager/models"
	"github.com/thiepwong/resident-manager/repositories"
)

type EmployeeService interface {
	Register(string, string, string, string) *models.Employee
	GetById(string) *models.Employee
	GetAll() *[]models.Employee
	Update(string, string, string, string) *models.Employee
}

type employeeServiceImp struct {
	employeeRepo repositories.EmployeeRepository
}

func NewEmployeeService(repo repositories.EmployeeRepository) EmployeeService {
	return &employeeServiceImp{
		employeeRepo: repo,
	}
}

func (s *employeeServiceImp) Register(departmentId string, name string, mobile string, accountId string) *models.Employee {
	var _emp = &models.Employee{
		ID:           "HoilamgiXXX",
		AccountId:    accountId,
		Name:         name,
		Mobile:       mobile,
		DepartmentId: departmentId,
		Address:      "Khong co gi",
		Role:         1,
		Status:       2,
	}
	rs, e := s.employeeRepo.Register(_emp)
	if e != nil {
		return nil
	}
	return rs
}

func (s *employeeServiceImp) GetById(id string) *models.Employee {
	rs, e := s.employeeRepo.GetById(id)
	if e != nil {
		return nil
	}
	return rs
}

func (s *employeeServiceImp) GetAll() *[]models.Employee {

	rs, e := s.employeeRepo.GetAll()
	if e != nil {
		return nil
	}
	return rs

}

func (s *employeeServiceImp) Update(departmentId string, name string, mobile string, accountId string) *models.Employee {
	var _emp = &models.Employee{}
	return _emp
}
