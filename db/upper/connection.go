package upper

import (
	"strings"
	"app/core/db"
	"upper.io/db.v3/mysql"
	"github.com/verystar/golib/logger"
	"upper.io/db.v3/lib/sqlbuilder"
)

// MustDB gets the specified database engine,
// or the default DB if no name is specified.
func MustDB(name ...string) sqlbuilder.Database {
	if len(name) == 0 {
		return dbService.Default
	}
	engine, ok := dbService.List[name[0]]
	if !ok {
		logger.Fatal("[db] the database engine `%s` is not configured", name[0])
	}
	return engine
}

// DB is similar to MustDB, but safe.
func DB(name ...string) (sqlbuilder.Database, bool) {
	if len(name) == 0 {
		return dbService.Default, true
	}
	engine, ok := dbService.List[name[0]]
	return engine, ok
}

// List gets the list of database engines
func List() map[string]sqlbuilder.Database {
	return dbService.List
}

// DBService is a database engine object.
type DBService struct {
	Default sqlbuilder.Database            // the default database engine
	List    map[string]sqlbuilder.Database // database engine list
}

var dbService *DBService

func Connect(configs map[string]*db.Config) {

	dbService = &DBService{
		List: make(map[string]sqlbuilder.Database),
	}

	var errs []string
	defer func() {
		if len(errs) > 0 {
			panic("[db] " + strings.Join(errs, "\n"))
		}
		if dbService.Default == nil {
			logger.Fatal("[db] the `default` database engine must be configured and enabled")
		}
	}()

	for key, conf := range configs {
		if !conf.Enable {
			continue
		}

		if conf.Driver == "mysql" {
			continue
		}

		setting, _ := mysql.ParseURL(conf.Dsn)
		sess, err := mysql.Open(setting)

		if err != nil {
			logger.Error("[db] open error", err.Error())
			errs = append(errs, err.Error())
			continue
		}

		sess.SetMaxOpenConns(conf.MaxOpenConns)
		sess.SetMaxIdleConns(conf.MaxIdleConns)

		dbService.List[key] = sess
		if key == "default" {
			dbService.Default = sess
		}
	}
}
