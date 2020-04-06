package database

import (
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"

	"github.com/Labbs/gokit/cfg"
)

// InitEngine - init xorm engine database
func InitEngine(engine string, dataSource []string, tables []interface{}, debug bool) *xorm.EngineGroup {
	cfg.Logger.Info("create database engine")
	e, err := xorm.NewEngineGroup(engine, dataSource)
	if err != nil {
		cfg.Logger.Fatal(err.Error())
	}

	e.ShowSQL(debug)
	e.ShowExecTime(debug)
	cfg.Logger.Info("create table")
	err = e.Sync2(tables...)
	if err != nil {
		cfg.Logger.Fatal(err.Error())
	}

	return e
}
