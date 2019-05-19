package repositories

import (
	"github.com/go-pg/pg"
	"github.com/thiepwong/resident-manager/models"
)

type RoomRepository interface {
	Add(*models.Room) (*models.Room, error)
	GetById(string) (*models.RoomModel, error)
	GetPagination(string, string, int, int, string) (*[]models.RoomModel, error)
	Update(*models.Room) (*models.Room, error)
	Delete(*models.Room) (bool, error)
}

type roomRepositoryContext struct {
	db *pg.DB
}

func NewRoomRepository(db *pg.DB) *roomRepositoryContext {
	return &roomRepositoryContext{
		db: db,
	}
}

func (r *roomRepositoryContext) Add(m *models.Room) (*models.Room, error) {
	return m, r.db.Insert(m)
}

func (r *roomRepositoryContext) GetById(id string) (*models.RoomModel, error) {

	var _room models.RoomModel
	r.db.Model(&_room).Column("room.*", "Block").Where("room.id=?", id).Select()
	var _side models.Side
	r.db.Model(&_side).Where("id=?", _room.SideId).Select()
	_room.Block.Side = _side
	if _room.Id == "" {
		return nil, nil
	}
	return &_room, nil

}

func (r *roomRepositoryContext) GetPagination(sideId string, blockId string, offset int, limit int, orderBy string) (*[]models.RoomModel, error) {
	var _room []models.RoomModel
	if orderBy == "" {
		orderBy = "id DESC"
	}
	var _side models.Side

	r.db.Model(&_room).Column("room.*", "Block").Where("block_id=?", blockId).Order(orderBy).Limit(limit).Offset(offset).Select()
	r.db.Model(&_side).Where("id=?", sideId).Select()

	for i := 0; i < len(_room); i++ {
		_room[i].Block.Side = _side
	}

	return &_room, nil
}

func (r *roomRepositoryContext) Update(m *models.Room) (*models.Room, error) {
	return m, r.db.Update(m)
}

func (r *roomRepositoryContext) Delete(m *models.Room) (bool, error) {
	return true, r.db.Delete(m)
}
