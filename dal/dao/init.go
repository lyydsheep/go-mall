package dao

import (
	"github.com/faiz/go-mall/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	dbMaster *gorm.DB
	dbSlave  *gorm.DB
)

func DB() *gorm.DB {
	return dbSlave
}

func init() {
	dbMaster = initDB(&config.DB.Master)
	dbSlave = initDB(&config.DB.Slave)
}

func initDB(option *config.DBConfigOptions) *gorm.DB {
	// 默认使用 MySQL
	db, err := gorm.Open(mysql.Open(option.Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDb, _ := db.DB()
	sqlDb.SetMaxOpenConns(option.MaxOpen)
	sqlDb.SetMaxIdleConns(option.MaxIdle)
	sqlDb.SetConnMaxLifetime(time.Duration(option.MaxLifeTime) * time.Second)
	return db
}
