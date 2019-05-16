package services

import (
	"errors"

	uuid "github.com/satori/go.uuid"
	"github.com/thiepwong/resident-manager/models"
	"github.com/thiepwong/resident-manager/repositories"
)

type BlockService interface {
	Add(string, string) (*models.Block, error)
	GetById(string) (*models.BlockModel, error)
	GetList(string, int, int, string) (*[]models.BlockModel, error)
	Update(string, string, string) (*models.Block, error)
	Delete(string) (bool, error)
}

type blockServiceImp struct {
	blockRepo repositories.BlockRepository
}

func NewBlockService(repo repositories.BlockRepository) BlockService {
	return &blockServiceImp{
		blockRepo: repo,
	}
}

func (s *blockServiceImp) Add(name string, sideId string) (*models.Block, error) {

	_id := uuid.Must(uuid.NewV4())
	_block := models.Block{
		Id:     _id.String(),
		Name:   name,
		SideId: sideId,
	}

	return s.blockRepo.Add(&_block)
}

func (s *blockServiceImp) GetById(id string) (*models.BlockModel, error) {
	return s.blockRepo.GetById(id)
}

func (s *blockServiceImp) GetList(sideId string, pageIndex int, pageSize int, orderBy string) (*[]models.BlockModel, error) {
	if pageIndex < 1 || pageSize < 1 {
		return nil, errors.New("Page index or page Size is invalid! Please check!")
	}
	var offset int
	offset = (pageIndex - 1) * pageSize
	rs, e := s.blockRepo.GetPagination(sideId, offset, pageSize, orderBy)
	if e != nil {
		return nil, e
	}
	return rs, nil
}

func (s *blockServiceImp) Update(id string, name string, sideId string) (*models.Block, error) {
	_side := &models.Block{
		Id:     id,
		Name:   name,
		SideId: sideId,
	}
	return s.blockRepo.Update(_side)
}

func (s *blockServiceImp) Delete(id string) (bool, error) {
	_side := &models.Block{Id: id}
	return s.blockRepo.Delete(_side)
}
