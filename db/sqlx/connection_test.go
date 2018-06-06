package sqlx

import (
	"os"
	"testing"

	"github.com/verystar/golib/config"
)

func TestConnect(t *testing.T) {
	configs := make(map[string]*config.Database)

	dsn := os.Getenv("MYSQL_TEST_DSN")

	if dsn == "" {
		dsn = "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Asia%2FShanghai"
	}

	configs["default"] = &config.Database{
		Enable: true,
		Driver: "mysql",
		Dsn:    dsn,
	}

	Connect(configs)
	link := DB()

	if link.DriverName() != "mysql" {
		t.Error("sqlx database connection error")
	}
}