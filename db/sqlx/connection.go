package sqlx

import (
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/verystar/golib/db"
)

var dbService = make(map[string]*sqlx.DB, 0)

// DB gets the specified database engine,
// or the default DB if no name is specified.
func DB(name ...string) *sqlx.DB {
	if len(name) == 0 {
		return dbService["default"]
	}
	engine, ok := dbService[name[0]]
	if !ok {
		log.Fatalf("[db] the database link `%s` is not configured", name[0])
	}
	return engine
}

// List gets the list of database engines
func List() map[string]*sqlx.DB {
	return dbService
}

func Connect(configs map[string]*db.Config) {

	var errs []string
	defer func() {
		if len(errs) > 0 {
			panic("[db] " + strings.Join(errs, "\n"))
		}

		if _, ok := dbService["default"]; !ok {
			log.Fatal("[db] the `default` database engine must be configured and enabled")
		}
	}()

	for key, conf := range configs {
		if !conf.Enable {
			continue
		}

		sess, err := sqlx.Connect(conf.Driver, conf.Dsn)

		if err != nil {
			errs = append(errs, err.Error())
			continue
		}
		log.Println("[db] connect:" + key)

		if conf.ShowSql {
			//sess.SetLogging(true)
		}

		sess.SetMaxOpenConns(conf.MaxOpenConns)
		sess.SetMaxIdleConns(conf.MaxIdleConns)

		dbService[key] = sess
	}
}
