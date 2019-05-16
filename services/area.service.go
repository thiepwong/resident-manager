package services

import (
	"errors"

	uuid "github.com/satori/go.uuid"
	"github.com/thiepwong/resident-manager/models"
	"github.com/thiepwong/resident-manager/repositories"
)

type AreaService interface {
	Add(string, string) (*models.Area, error)
	GetById(string) (*models.AreaModel, error)
	GetList(string, int, int, string) (*[]models.AreaModel, error)
	Update(string, string, string) (*models.Area, error)
	Delete(string) (bool, error)
}

type areaServiceImp struct {
	areaRepo repositories.AreaRepository
}

func NewAreaService(repo repositories.AreaRepository) AreaService {
	return &areaServiceImp{
		areaRepo: repo,
	}
}

func (s *areaServiceImp) Add(name string, sideId string) (*models.Area, error) {

	_id := uuid.Must(uuid.NewV4())
	_area := models.Area{
		Id:     _id.String(),
		Name:   name,
		SideId: sideId,
	}

	return s.areaRepo.Add(&_area)
}

func (s *areaServiceImp) GetById(id string) (*models.AreaModel, error) {
	return s.areaRepo.GetById(id)
}

func (s *areaServiceImp) GetList(sideId string, pageIndex int, pageSize int, orderBy string) (*[]models.AreaModel, error) {
	if pageIndex < 1 || pageSize < 1 {
		return nil, errors.New("Page index or page Size is invalid! Please check!")
	}
	var offset int
	offset = (pageIndex - 1) * pageSize
	rs, e := s.areaRepo.GetPagination(sideId, offset, pageSize, orderBy)
	if e != nil {
		return nil, e
	}
	return rs, nil
}

func (s *areaServiceImp) Update(id string, name string, sideId string) (*models.Area, error) {
	_side := &models.Area{
		Id:     id,
		Name:   name,
		SideId: sideId,
	}
	return s.areaRepo.Update(_side)
}

func (s *areaServiceImp) Delete(id string) (bool, error) {
	_side := &models.Area{Id: id}
	return s.areaRepo.Delete(_side)
}
