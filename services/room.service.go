package services

import (
	"errors"

	uuid "github.com/satori/go.uuid"
	"github.com/thiepwong/resident-manager/models"
	"github.com/thiepwong/resident-manager/repositories"
)

type RoomService interface {
	Add(string, string, string) (*models.Room, error)
	GetById(string) (*models.RoomModel, error)
	GetList(string, string, int, int, string) (*[]models.RoomModel, error)
	Update(string, string, string, string) (*models.Room, error)
	Delete(string) (bool, error)
}

type roomServiceImp struct {
	roomRepo repositories.RoomRepository
}

func NewRoomService(repo repositories.RoomRepository) RoomService {
	return &roomServiceImp{
		roomRepo: repo,
	}
}

func (s *roomServiceImp) Add(roomNo string, sideId string, blockId string) (*models.Room, error) {

	_id := uuid.Must(uuid.NewV4())
	_room := models.Room{
		Id:      _id.String(),
		RoomNo:  roomNo,
		SideId:  sideId,
		BlockId: blockId,
	}

	return s.roomRepo.Add(&_room)
}

func (s *roomServiceImp) GetById(id string) (*models.RoomModel, error) {
	return s.roomRepo.GetById(id)
}

func (s *roomServiceImp) GetList(sideId string, blockId string, pageIndex int, pageSize int, orderBy string) (*[]models.RoomModel, error) {
	if pageIndex < 1 || pageSize < 1 {
		return nil, errors.New("Page index or page Size is invalid! Please check!")
	}
	var offset int
	offset = (pageIndex - 1) * pageSize
	rs, e := s.roomRepo.GetPagination(sideId, blockId, offset, pageSize, orderBy)
	if e != nil {
		return nil, e
	}
	return rs, nil
}

func (s *roomServiceImp) Update(id string, roomNo string, sideId string, blockId string) (*models.Room, error) {
	_room := &models.Room{
		Id:      id,
		RoomNo:  roomNo,
		SideId:  sideId,
		BlockId: blockId,
	}
	return s.roomRepo.Update(_room)
}

func (s *roomServiceImp) Delete(id string) (bool, error) {
	_room := &models.Room{Id: id}
	return s.roomRepo.Delete(_room)
}
