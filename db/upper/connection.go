package upper

import (
	"strings"

	"github.com/verystar/golib/cache"
	"github.com/verystar/golib/db"
	"github.com/verystar/golib/logger"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/mysql"
)

type Database struct {
	Link          sqlbuilder.Database
	CachedColumns *cache.Cache
}

var dbService map[string]*Database

// MustDB gets the specified database engine,
// or the default DB if no name is specified.
func MustDB(name ...string) *Database {
	if len(name) == 0 {
		return dbService["default"]
	}
	engine, ok := dbService[name[0]]
	if !ok {
		logger.Fatal("[db] the database link `%s` is not configured", name[0])
	}
	return engine
}

// List gets the list of database engines
func List() map[string]*Database {
	return dbService
}

func Connect(configs map[string]*db.Config) {

	var errs []string
	defer func() {
		if len(errs) > 0 {
			panic("[db] " + strings.Join(errs, "\n"))
		}

		if _, ok := dbService["default"]; !ok {
			logger.Fatal("[db] the `default` database engine must be configured and enabled")
		}
	}()

	for key, conf := range configs {
		if !conf.Enable {
			continue
		}

		if conf.Driver != "mysql" {
			continue
		}

		setting, _ := mysql.ParseURL(conf.Dsn)
		sess, err := mysql.Open(setting)

		if err != nil {
			logger.Error("[db] open error", err.Error())
			errs = append(errs, err.Error())
			continue
		}
		logger.Debug("[db] connect:" + key)

		if conf.ShowSql {
			sess.SetLogging(true)
		}

		sess.SetMaxOpenConns(conf.MaxOpenConns)
		sess.SetMaxIdleConns(conf.MaxIdleConns)

		link := &Database{
			Link:          sess,
			CachedColumns: cache.NewCache(),
		}

		dbService[key] = link
	}
}
