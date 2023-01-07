package db

import (
	"xorm.io/xorm"
)

var Engine *xorm.Engine

func init() {
	var err error
	Engine, err = xorm.NewEngine("sqlite3", "./data.db")
	if err != nil {
		panic("xorm init failed")
	}
}
