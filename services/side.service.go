package services

import (
	"errors"

	uuid "github.com/satori/go.uuid"
	"github.com/thiepwong/resident-manager/models"
	"github.com/thiepwong/resident-manager/repositories"
)

type SideService interface {
	Add(string, string, string, string, string) (*models.Side, error)
	GetById(string) (*models.Side, error)
	GetList(int, int, string) (*[]models.Side, error)
	Update(string, string, string, string, string, string) (*models.Side, error)
	Delete(string) (bool, error)
}

type sideService struct {
	sideRepo repositories.SideRepository
}

func NewSideService(repo repositories.SideRepository) SideService {
	return &sideService{
		sideRepo: repo,
	}
}

func (s *sideService) Add(name string, address string, ip string, cover string, hotline string) (*models.Side, error) {

	_id := uuid.Must(uuid.NewV4())
	_side := models.Side{
		Id:      _id.String(),
		Name:    name,
		Address: address,
		Ip:      ip,
		Cover:   cover,
		Hotline: hotline,
	}
	return s.sideRepo.Add(&_side)
}

func (s *sideService) GetById(id string) (*models.Side, error) {
	return s.sideRepo.GetById(id)
}

func (s *sideService) GetList(pageIndex int, pageSize int, orderBy string) (*[]models.Side, error) {
	if pageIndex < 1 || pageSize < 1 {
		return nil, errors.New("Page index or page Size is invalid! Please check!")
	}
	var offset int
	offset = (pageIndex - 1) * pageSize
	rs, e := s.sideRepo.GetPagination(offset, pageSize, orderBy)
	if e != nil {
		return nil, e
	}
	return rs, nil
}

func (s *sideService) Update(id string, name string, address string, ip string, cover string, hotline string) (*models.Side, error) {
	_side := &models.Side{
		Id:      id,
		Name:    name,
		Address: address,
		Ip:      ip,
		Cover:   cover,
		Hotline: hotline,
	}
	return s.sideRepo.Update(_side)
}

func (s *sideService) Delete(id string) (bool, error) {
	_side := &models.Side{Id: id}
	return s.sideRepo.Delete(_side)
}
