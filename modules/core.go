package modules

import (
	"database/sql"
	"fmt"

	"github.com/tututuf/letsStudy_goServer/env"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB
var DbErr error
var Fdb *sql.DB

func init() {
	user, _ := env.Config.GetValue("mysql", "user")
	password, _ := env.Config.GetValue("mysql", "password")
	database, _ := env.Config.GetValue("mysql", "database")
	url, _ := env.Config.GetValue("mysql", "url")
	port, _ := env.Config.GetValue("mysql", "port")
	origin := url + ":" + port
	dsn := user + ":" + password + "@tcp(" + origin + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, DbErr = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 单数表名
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	Fdb, _ = DB.DB()
	if DbErr != nil {
		fmt.Println(DbErr)
	}
}
