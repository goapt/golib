package redis

import (
	"github.com/go-redis/redis"
	"github.com/verystar/golib/logger"
	"strings"
)

var (
	redisList map[string]*redis.Client
	errs      []string
)

type Config struct {
	Server   string
	Password string
	DB       int
}

func Client(name ... string) (*redis.Client, bool) {
	key := "default"
	if len(name) > 0 {
		key = name[0]
	}
	pool, ok := redisList[key]
	return pool, ok
}

func Connect(configs map[string]Config) {
	defer func() {
		if len(errs) > 0 {
			panic("[redis] " + strings.Join(errs, "\n"))
		}
	}()

	redisList = make(map[string]*redis.Client)
	for name, conf := range configs {
		r := newRedis(&conf)
		logger.Debug("[redis] connect:" + conf.Server)

		_, err := r.Ping().Result()
		if err != nil {
			logger.Error("[redis] ping", err.Error())
			errs = append(errs, err.Error())
		}

		redisList[name] = newRedis(&conf)
	}
}

// 创建 redis pool
func newRedis(conf *Config) *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     conf.Server,
		Password: conf.Password, // no password set
		DB:       conf.DB,       // use default DB
	})
	return client
}