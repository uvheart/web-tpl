package app

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"web_tpl1/app/core/config"
	"web_tpl1/app/core/db"
	redis1 "web_tpl1/app/core/redis"
)

var Config config.Model

func Init(prjHome string) error {

	return Config.LoadConfig(prjHome)

}

func DBW(keys ...string) *gorm.DB {

	k := "default"
	if len(keys) > 0 {
		k = keys[0]
	}
	conf, ok := Config.DB[k]

	if !ok {
		panic(fmt.Sprintf("db config %s not found", k))
	}

	cacheKey := fmt.Sprintf("%s_write", k)
	return db.Load(conf.Write, cacheKey)

}

func RedisW(keys ...string) *redis.Client {

	k := "default"
	if len(keys) > 0 {
		k = keys[0]
	}
	conf, ok := Config.Redis[k]

	if !ok {
		panic(fmt.Sprintf("redis config %s not found", k))
	}

	fmt.Println(conf)
	//cacheKey := fmt.Sprintf("redis_%s_write", k)
	return redis1.InitRedis(false, conf)

}

func DBR(keys ...string) *gorm.DB {

	k := "default"
	if len(keys) > 0 {
		k = keys[0]
	}
	conf, ok := Config.DB[k]

	if !ok {
		panic(fmt.Sprintf("db config %s not found", k))
	}

	fmt.Println(conf)
	return nil

	cacheKey := fmt.Sprintf("%s_read", k)
	return db.Load(conf.Read, cacheKey)

}
