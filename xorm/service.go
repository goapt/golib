package xorm

import (
	"strings"
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql" //mysql
	"github.com/verystar/golib/logger"
)

// DBService is a database engine object.
type DBService struct {
	Default *xorm.Engine            // the default database engine
	List    map[string]*xorm.Engine // database engine list
}

var dbService *DBService

type Config struct {
	Enable       bool
	Driver       string
	Dsn          string
	MaxOpenConns int  `toml:"max_open_conns"`
	MaxIdleConns int  `toml:"max_idle_conns"`
	Cache        bool
	ShowExecTime bool `toml:"show_exec_time"`
	ShowSql      bool `toml:"show_sql"`
}

func Connect(configs map[string]*Config) {

	dbService = &DBService{
		List: map[string]*xorm.Engine{},
	}
	var errs []string
	defer func() {
		if len(errs) > 0 {
			panic("[xorm] " + strings.Join(errs, "\n"))
		}
		if dbService.Default == nil {
			logger.Fatal("[xorm] the `default` database engine must be configured and enabled")
		}
	}()

	for key, conf := range configs {
		if !conf.Enable {
			continue
		}
		engine, err := xorm.NewEngine(conf.Driver, conf.Dsn)
		if err != nil {
			logger.Error("[xorm] new", err.Error())
			errs = append(errs, err.Error())
			continue
		}
		err = engine.Ping()
		if err != nil {
			logger.Error("[xorm] ping", err.Error())
			errs = append(errs, err.Error())
			continue
		}
		engine.SetMaxOpenConns(conf.MaxOpenConns)
		engine.SetMaxIdleConns(conf.MaxIdleConns)
		engine.SetDisableGlobalCache(conf.Cache)
		engine.ShowSQL(conf.ShowSql)
		engine.ShowExecTime(conf.ShowExecTime)

		dbService.List[key] = engine
		if key == "default" {
			dbService.Default = engine
		}
	}
}