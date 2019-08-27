package repositories

import (
	"fmt"

	"github.com/go-pg/pg"
	"github.com/thiepwong/resident-manager/models"
)

type RoomRepository interface {
	Add(*models.Room) (*models.Room, error)
	GetById(string) (*models.RoomModel, error)
	GetPagination(string, string, string, int, int, string) (interface{}, error)
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

func (r *roomRepositoryContext) GetPagination(sideId string, blockId string, roomName string, offset int, limit int, orderBy string) (interface{}, error) {
	var _room []models.RoomQueyModel
	var _result models.ModelResult
	var count int
	if orderBy == "" {
		orderBy = "room_no ASC"
	}

	count_query := fmt.Sprintf("select count(r.id) from resident.room r left join resident.block rb on rb.id = r.block_id left join resident.side rs on rs.id = r.side_id where  r.side_id='%s' AND (r.block_id in (select id from resident.block b where LOWER(b.id) like LOWER('%s')) or '%s' = '') AND ( lower(room_no) like '%s' or  '%s' = '' )", sideId, "%"+blockId+"%", blockId, "%"+roomName+"%", roomName)

	select_query := fmt.Sprintf("select r.id, r.room_no, rb.id blockid, rb.name blockname, rs.id sideid, rs.name sidename from resident.room r left join resident.block rb on rb.id = r.block_id left join resident.side rs on rs.id = r.side_id where  r.side_id='%s' AND (r.block_id in (select id from resident.block b where LOWER(b.id) like LOWER('%s')) or '%s' = '') AND ( lower(room_no) like '%s' or  '%s' = '' ) ORDER BY %s offset %d limit %d ", sideId, "%"+blockId+"%", blockId, "%"+roomName+"%", roomName, orderBy, offset, limit)

	_, e := r.db.Query(&count, count_query)

	_, e = r.db.Query(&_room, select_query)

	if e != nil {
		return nil, e
	}

	_result.TotalRecord = count
	_result.Rows = _room
	return _result, nil
}

func (r *roomRepositoryContext) Update(m *models.Room) (*models.Room, error) {
	return m, r.db.Update(m)
}

func (r *roomRepositoryContext) Delete(m *models.Room) (bool, error) {
	return true, r.db.Delete(m)
}
