package services

import (
	"errors"

	uuid "github.com/satori/go.uuid"
	"github.com/thiepwong/resident-manager/models"
	"github.com/thiepwong/resident-manager/repositories"
)

type DepartmentService interface {
	Add(*models.Department) (*models.Department, error)
	GetById(string) (*models.Department, error)
	GetList(int, int, string) (*[]models.Department, error)
	Update(string, string, string) (*models.Department, error)
	Delete(string) (bool, error)
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
	if m.Name == "" {
		return nil, errors.New("invalid infomation!")
	}

	_id := uuid.Must(uuid.NewV4())
	m.ID = _id.String()
	return s.departmentRepo.Add(m)

}

func (s *departmentServiceImp) GetById(id string) (*models.Department, error) {
	return s.departmentRepo.GetById(id)
}

func (s *departmentServiceImp) GetList(pageIndex int, pageSize int, orderBy string) (*[]models.Department, error) {
	if pageIndex < 1 || pageSize < 1 {
		return nil, errors.New("Page index or page Size is invalid! Please check!")
	}
	var offset int
	offset = (pageIndex - 1) * pageSize
	rs, e := s.departmentRepo.GetPagination(offset, pageSize, orderBy)
	if e != nil {
		return nil, e
	}
	return rs, nil
}

func (s *departmentServiceImp) Update(id string, name string, sideId string) (*models.Department, error) {
	var _dept = models.Department{ID: id, Name: name, SideId: sideId}
	return s.departmentRepo.Update(&_dept)
}

func (s *departmentServiceImp) Delete(id string) (bool, error) {
	var _dept = models.Department{ID: id}
	return s.departmentRepo.Delete(&_dept)
}