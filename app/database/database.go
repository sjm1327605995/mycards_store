package database

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db *gorm.DB
)

func GetDB() *gorm.DB {
	return db
}

func InitDB() {
	var err error
	host := viper.GetString("postgres")
	db, err = gorm.Open(postgres.Open(host), &gorm.Config{
		CreateBatchSize:                          1000,
		Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(fmt.Errorf("数据库连接失败,%s", err.Error()))
	}

}
