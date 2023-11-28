package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
	"web_tpl1/app/core/config"
)

const (
	dbTimeout      = 5000
	dbWriteTimeout = 5000
	dbReadTimeout  = 5000
)

var dbInstance = make(map[string]*gorm.DB)
var dbLocker sync.RWMutex

func getDbInstance(conf config.DBItemConf) *gorm.DB {

	timeout := dbTimeout
	if conf.Timeout > 0 {
		timeout = conf.Timeout
	}

	writeTimeout := dbWriteTimeout
	if conf.Timeout > 0 {
		writeTimeout = conf.Timeout
	}
	readTimeout := dbReadTimeout
	if conf.Timeout > 0 {
		readTimeout = conf.Timeout
	}

	dsnConf := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local&timeout=%dms&writeTimeout=%dms&readTimeout=%dms",
		conf.Username,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database,
		conf.Charset,
		timeout,
		writeTimeout,
		readTimeout,
	)
	//dsnConf := "root:123456@tcp(127.0.0.1:3306)/kbh?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsnConf), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}

	//数据库池
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(conf.MaxIdleConns)
	sqlDB.SetMaxOpenConns(conf.MaxOpenConns)

	return db

}

func Load(conf config.DBItemConf, key string) *gorm.DB {

	dbLocker.RLock()
	db, ok := dbInstance[key]

	if ok {
		dbLocker.RUnlock()
		return db
	}
	dbLocker.RUnlock()
	dbLocker.Lock()
	defer dbLocker.Unlock()

	if _, exist := dbInstance[key]; exist {
		return dbInstance[key]
	}

	dbInstance[key] = getDbInstance(conf)

	return dbInstance[key]
}
