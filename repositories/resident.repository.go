package repositories

import (
	"github.com/go-pg/pg"
	"github.com/thiepwong/resident-manager/models"
)

type ResidentRepository interface {
	GetPagination(string, string, int, int, string) (*[]models.ResidentRoom, error)
	GetById(string) (*models.ResidentRoom, error)
}

type residentRepositoryContext struct {
	db *pg.DB
}

func NewResidentRepository(db *pg.DB) *residentRepositoryContext {
	return &residentRepositoryContext{
		db: db,
	}
}

func (r *residentRepositoryContext) GetPagination(sideId string, search string, offset int, limit int, pageOrder string) (*[]models.ResidentRoom, error) {

	if pageOrder == "" {
		pageOrder = "id ASC"
	}

	var _room []models.Room
	var _resident []models.ResidentRoom

	r.db.Model(&_room).Column("room.id").Where("room.side_id=?", sideId).Select()
	var _r []string
	for i := 0; i < len(_room); i++ {
		_r = append(_r, _room[i].Id)
	}
	//var levels = []string{"ROOM0000001", "ROOM0000002", "ROOM0000003"}
	r.db.Model(&_resident).Column("resident_room_mapping.*", "Resident", "Room", "Room.Block.Side", "Room.Side", "Room.Block").Where("resident_room_mapping.room_id in (?)", pg.In(_r)).Order(pageOrder).Limit(limit).Offset(offset).Select()

	return &_resident, nil
}

func (r *residentRepositoryContext) GetById(id string) (*models.ResidentRoom, error) {
	var _resident models.ResidentRoom
	r.db.Model(&_resident).Column("resident_room_mapping.*", "Resident", "Room", "Room.Block.Side", "Room.Side", "Room.Block").Where("resident_room_mapping.resident_id=?", id).Select()
	return &_resident, nil
}
