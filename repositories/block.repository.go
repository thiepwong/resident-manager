package repositories

import (
	"github.com/go-pg/pg"
	"github.com/thiepwong/resident-manager/models"
)

type BlockRepository interface {
	Add(*models.Block) (*models.Block, error)
	GetById(string) (*models.BlockModel, error)
	GetPagination(int, int, string) (*[]models.BlockModel, error)
	Update(*models.Block) (*models.Block, error)
	Delete(*models.Block) (bool, error)
}

type blockRepositoryContext struct {
	db *pg.DB
}

func NewBlockRepository(db *pg.DB) *blockRepositoryContext {
	return &blockRepositoryContext{
		db: db,
	}
}

func (r *blockRepositoryContext) Add(m *models.Block) (*models.Block, error) {
	return m, r.db.Insert(m)
}

func (r *blockRepositoryContext) GetById(id string) (*models.BlockModel, error) {
	var _block models.BlockModel
	return &_block, r.db.Model(&_block).Where("id=?", id).Select()
}

func (r *blockRepositoryContext) GetPagination(offset int, limit int, orderBy string) (*[]models.BlockModel, error) {
	var _block []models.BlockModel
	if orderBy == "" {
		orderBy = "id DESC"
	}
	r.db.Model(&_block).Order(orderBy).Limit(limit).Offset(offset).Select()
	return &_block, nil
}

func (r *blockRepositoryContext) Update(m *models.Block) (*models.Block, error) {
	return m, r.db.Update(m)
}

func (r *blockRepositoryContext) Delete(m *models.Block) (bool, error) {
	return true, r.db.Delete(m)
}
