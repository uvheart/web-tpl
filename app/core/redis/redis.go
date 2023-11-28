package redis

import (
	"github.com/redis/go-redis/v9"
	"sync"
	"web_tpl1/app/core/config"
)

const (
	dbTimeout      = 5000
	dbWriteTimeout = 5000
	dbReadTimeout  = 5000
)

// var dbInstance = make(map[string]*redis.Client)
var dbLocker sync.RWMutex

func InitRedis(isWrite bool, conf config.Redis) *redis.Client {

	var redisConf config.RedisItem

	if isWrite {
		redisConf = conf.Write
	} else {
		redisConf = conf.Read
	}
	return redis.NewClient(&redis.Options{
		Addr:         redisConf.Addr,
		PoolSize:     redisConf.PoolSize,
		Password:     redisConf.Password,
		ReadTimeout:  redisConf.ReadTimeout,
		WriteTimeout: redisConf.WriteTimeOut,
		MaxRetries:   redisConf.Retries,
		MinIdleConns: redisConf.MinIdleConns,
		DB:           redisConf.DB,
	})

}
