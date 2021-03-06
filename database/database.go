package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

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

// InitEngineSQLite - init xorm engine database for sqlite3
func InitEngineSQLite(path string, tables []interface{}, debug bool) *xorm.Engine {
	cfg.Logger.Info("create database engine")
	e, err := xorm.NewEngine("sqlite3", path)
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
