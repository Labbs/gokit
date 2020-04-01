package database

import (
	"github.com/go-xorm/xorm"

	"github.com/Labbs/gokit/cfg"
)

func InitEngine(engine string, dataSource []string, table ...interface{}) *xorm.EngineGroup {
	cfg.Logger.Info("create database engine")
	engine, err := xorm.NewEngineGroup(cfg.Config["database_engine"].(string), dataSourceName)
	if err != nil {
		cfg.Logger.Fatal(err.Error())
	}

	cfg.Logger.Info("create table")
	err = engine.Sync2(table)
	if err != nil {
		cfg.Logger.Fatal(err.Error())
	}
}
