package repositories

import (
	"errors"
	"strings"

	"github.com/go-pg/pg"
	"github.com/thiepwong/resident-manager/models"
)

type ResidentRepository interface {
	GetPagination(string, string, string, string, string, string, string, int, int, string) (*[]models.ResidentRoom, error)
	GetById(string) (*models.ResidentRoom, error)
	Add(*models.Resident, *models.ResidentRoom) (*models.ResidentRoom, error)
	Update(*models.Resident, *models.ResidentRoom) (*models.ResidentRoom, error)
	Delete(residentRoomId string) (bool, error)
}

type residentRepositoryContext struct {
	db *pg.DB
}

func NewResidentRepository(db *pg.DB) *residentRepositoryContext {
	return &residentRepositoryContext{
		db: db,
	}
}

func (r *residentRepositoryContext) GetPagination(sideId string, search string, block string, room string, name string, mobile string, email string, offset int, limit int, pageOrder string) (*[]models.ResidentRoom, error) {

	if pageOrder == "" {
		pageOrder = "id ASC"
	}

	var _room []models.RoomModel
	var _resident []models.ResidentRoom
	if room == "" {
		r.db.Model(&_room).Column("room.*", "Block").Where("room.side_id=?", sideId).Select()

	} else {
		r.db.Model(&_room).Column("room.*", "Block").Where("room.side_id=?", sideId).Where("LOWER(room.room_no) LIKE ?", "%"+strings.ToLower(room)+"%").Select()

	}
	var _r []string
	for i := 0; i < len(_room); i++ {
		if block == "" {
			_r = append(_r, _room[i].Id)
		} else {
			if strings.Contains(strings.ToLower(_room[i].Block.Name), strings.ToLower(block)) == true {
				_r = append(_r, _room[i].Id)
			}
		}

	}

	var _fullname []models.Resident
	// var _mobile []models.Resident
	// var _email []models.Resident
	if name != "" {
		r.db.Model(&_fullname).Column("resident.*").Where("LOWER(resident.full_name) LIKE ?", "%"+strings.ToLower(name)+"%").WhereOr("LOWER(resident.phone_no) LIKE ?", "%"+strings.ToLower(name)+"%").WhereOr("LOWER(resident.email) LIKE ?", "%"+strings.ToLower(name)+"%").Select()
	}

	// if mobile != "" {
	// 	r.db.Model(&_mobile).Column("resident.*").Where("resident.phone_no LIKE ?", "%"+mobile+"%").Select()
	// }

	// if email != "" {
	// 	r.db.Model(&_email).Column("resident.*").Where("resident.email LIKE ?", "%"+email+"%").Select()
	// }

	// _fullname = append(_fullname, _mobile...)
	// _fullname = append(_fullname, _email...)

	var _rid []string
	for i := 0; i < len(_fullname); i++ {
		_rid = append(_rid, _fullname[i].Id)
	}

	if len(_rid) > 0 {
		r.db.Model(&_resident).Column("resident_room_mapping.*", "Resident", "Room", "Room.Block.Side", "Room.Side", "Room.Block").Where("resident_room_mapping.room_id in (?)", pg.In(_r)).Where("resident_room_mapping.deleted <>?", true).Where("resident_room_mapping.resident_id in (?)", pg.In(_rid)).Order(pageOrder).Limit(limit).Offset(offset).Select()

	} else {
		r.db.Model(&_resident).Column("resident_room_mapping.*", "Resident", "Room", "Room.Block.Side", "Room.Side", "Room.Block").Where("resident_room_mapping.room_id in (?)", pg.In(_r)).Where("resident_room_mapping.deleted <>?", true).Order(pageOrder).Limit(limit).Offset(offset).Select()
	}

	return &_resident, nil
}

func (r *residentRepositoryContext) GetById(id string) (*models.ResidentRoom, error) {
	var _resident models.ResidentRoom
	r.db.Model(&_resident).Column("resident_room_mapping.*", "Resident", "Room", "Room.Block.Side", "Room.Side", "Room.Block").Where("resident_room_mapping.resident_id=?", id).Where("resident_room_mapping.deleted <>?", true).Select()
	return &_resident, nil
}

func (r *residentRepositoryContext) Add(res *models.Resident, roo *models.ResidentRoom) (*models.ResidentRoom, error) {

	e := r.db.Insert(res)
	if e != nil {
		return nil, errors.New("Insert error!")
	}
	roo.ResidentId = res.Id
	e = r.db.Insert(roo)
	var _resident models.ResidentRoom
	r.db.Model(&_resident).Column("resident_room_mapping.*", "Resident", "Room", "Room.Block.Side", "Room.Side", "Room.Block").Where("resident_room_mapping.id=?", roo.Id).Select()
	return &_resident, e
}

func (r *residentRepositoryContext) Update(res *models.Resident, resR *models.ResidentRoom) (*models.ResidentRoom, error) {

	e := r.db.Update(res)
	if e != nil {
		return nil, errors.New("Update error!")
	}

	_r, err := r.db.Model(resR).Set("room_id = ?", resR.RoomId).Where("id = ?", resR.Id).Returning("*").Update()
	if err != nil {
		return nil, err

	}
	if _r != nil {

	}

	r.db.Model(resR).Column("resident_room_mapping.*", "Resident", "Room", "Room.Block.Side", "Room.Side", "Room.Block").Where("resident_room_mapping.id=?", resR.Id).Select()

	return resR, e
}

func (r residentRepositoryContext) Delete(resRId string) (bool, error) {
	var resR = models.ResidentRoom{}
	_r, err := r.db.Model(&resR).Set("deleted = ?", true).Where("id = ?", resRId).Returning("*").Update()
	if err != nil {
		return false, err

	}
	if _r != nil {

	}
	return true, nil
}
