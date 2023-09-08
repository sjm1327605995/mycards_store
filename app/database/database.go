package database

import (
	"fmt"
	"github.com/sjm1327605995/mycards_store/app/models"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strings"
)

var (
	db *gorm.DB
)

func GetDB() *gorm.DB {
	return db
}

func InitDB() {
	var (
		err       error
		dialector gorm.Dialector
	)
	switch strings.ToLower(viper.GetString("db.type")) {
	case "mysql":
		dialector = Mysql()
	case "postgres", "pg", "pgsql":
		dialector = Postgres()
	case "sqlite", "":
		dialector = SQLite()
	default:
		panic("不支持的数据库类型")
	}

	db, err = gorm.Open(dialector, &gorm.Config{
		CreateBatchSize:                          1000,
		Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(fmt.Errorf("数据库连接失败,%s", err.Error()))
	}

	if viper.GetBool("db.init") {
		_ = db.AutoMigrate(&models.Decks{})
	}
}

func Mysql() gorm.Dialector {

	var (
		host     = viper.GetString("db.host")
		user     = viper.GetString("db.user")
		dbname   = viper.GetString("db.name")
		password = viper.GetString("db.pwd")
		port     = viper.GetInt("db.port")
	)
	return mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname))

}
func Postgres() gorm.Dialector {

	var (
		host     = viper.GetString("db.host")
		user     = viper.GetString("db.user")
		dbname   = viper.GetString("db.name")
		password = viper.GetString("db.pwd")
		port     = viper.GetInt("db.port")
	)
	return postgres.Open(
		fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", host, user, password, dbname, port),
	)
}
func SQLite() gorm.Dialector {
	return sqlite.Open("mycards.db")
}
