package datasources

import (
	"strconv"

	"github.com/thiepwong/resident-manager/common"

	"github.com/go-pg/pg"
)

func GetPg(cfg common.CfgPg) (*pg.DB, error) {
	db := pg.Connect(&pg.Options{Addr: cfg.Host + ":" + strconv.Itoa(cfg.Port), User: cfg.Username, Password: cfg.Password, Database: cfg.DbName})
	return db, nil
}
