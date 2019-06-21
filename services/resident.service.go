package services

import (
	"errors"

	"github.com/thiepwong/resident-manager/models"
	"github.com/thiepwong/resident-manager/repositories"
)

type ResidentService interface {
	GetById(string) (*models.ResidentRoom, error)
	GetList(string, string, int, int, string) (*[]models.ResidentRoom, error)
}

type residentServiceImp struct {
	residentRepo repositories.ResidentRepository
}

func NewResidentService(repo repositories.ResidentRepository) ResidentService {
	return &residentServiceImp{
		residentRepo: repo,
	}
}

func (s *residentServiceImp) GetById(id string) (*models.ResidentRoom, error) {
	return s.residentRepo.GetById(id)
}

func (s *residentServiceImp) GetList(sideId string, search string, pageIndex int, pageSize int, pageOrder string) (*[]models.ResidentRoom, error) {
	if pageIndex < 1 || pageSize < 1 {
		return nil, errors.New("Page index or page Size is invalid! Please check!")
	}
	var offset int
	offset = (pageIndex - 1) * pageSize
	return s.residentRepo.GetPagination(sideId, search, offset, pageSize, pageOrder)
}
