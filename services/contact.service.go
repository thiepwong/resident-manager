package services

import (
	"errors"

	uuid "github.com/satori/go.uuid"
	"github.com/thiepwong/resident-manager/models"
	"github.com/thiepwong/resident-manager/repositories"
)

type ContactService interface {
	Add(string, string) (*models.Contact, error)
	GetById(string) (*models.ContactModel, error)
	GetList(string, int, int, string) (*[]models.ContactModel, error)
	Update(string, string, string, string) (*models.Contact, error)
	Delete(string) (bool, error)
}

type contactServiceImp struct {
	contactRepo repositories.ContactRepository
}

func NewContactService(repo repositories.ContactRepository) ContactService {
	return &contactServiceImp{
		contactRepo: repo,
	}
}

func (s *contactServiceImp) Add(name string, departmentId string) (*models.Contact, error) {

	_id := uuid.Must(uuid.NewV4())
	_contact := models.Contact{
		Id:           _id.String(),
		Name:         name,
		DepartmentId: departmentId,
	}

	return s.contactRepo.Add(&_contact)
}

func (s *contactServiceImp) GetById(id string) (*models.ContactModel, error) {
	return s.contactRepo.GetById(id)
}

func (s *contactServiceImp) GetList(departmentId string, pageIndex int, pageSize int, orderBy string) (*[]models.ContactModel, error) {
	if pageIndex < 1 || pageSize < 1 {
		return nil, errors.New("Page index or page Size is invalid! Please check!")
	}
	var offset int
	offset = (pageIndex - 1) * pageSize
	rs, e := s.contactRepo.GetPagination(departmentId, offset, pageSize, orderBy)
	if e != nil {
		return nil, e
	}
	return rs, nil
}

func (s *contactServiceImp) Update(id string, name string, departmentId string, phone string) (*models.Contact, error) {
	_side := &models.Contact{
		Id:           id,
		Name:         name,
		DepartmentId: departmentId,
		PhoneNumber:  phone,
	}
	return s.contactRepo.Update(_side)
}

func (s *contactServiceImp) Delete(id string) (bool, error) {
	_side := &models.Contact{Id: id}
	return s.contactRepo.Delete(_side)
}
