package services

import (
	"github.com/thiepwong/resident-manager/models"
	"github.com/thiepwong/resident-manager/repositories"
)

type DepartmentService interface {
	Add(*models.Department) (*models.Department, error)
	GetById(string) (*models.Department, error)
	GetList(int, int, string) (*[]models.Department, error)
}

type departmentServiceImp struct {
	departmentRepo repositories.DepartmentRepository
}

func NewDepartmentService(repo repositories.DepartmentRepository) DepartmentService {
	return &departmentServiceImp{

		departmentRepo: repo,
	}
}

func (s *departmentServiceImp) Add(m *models.Department) (*models.Department, error) {
	return m, nil
}

func (s *departmentServiceImp) GetById(id string) (*models.Department, error) {
	return s.departmentRepo.GetById(id)
}

func (s *departmentServiceImp) GetList(page int, size int, order string) (*[]models.Department, error) {
	var _departList []models.Department
	return &_departList, nil
}
