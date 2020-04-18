package driver

import (
	"fmt"
	"passport.xinfos.com/configs"
	"time"

	"github.com/jinzhu/gorm"

	_ "database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DB *gorm.DB
)

func InitDB() {
	initMySQL()
}

func initMySQL() {
	cfg := configs.Get()
	var err error
	timeout := "10s" //连接超时，10秒
	connect := fmt.Sprintf("%v?charset=utf8&parseTime=True&loc=Local&timeout=%v", cfg.Mysql.Uri, timeout)

	orm, err := gorm.Open("mysql", connect)

	if err != nil {
		panic(err)
	}
	orm.LogMode(true)

	orm.DB().SetConnMaxLifetime(time.Second * 600)
	orm.DB().SetMaxIdleConns(10)
	orm.DB().SetMaxOpenConns(200)

	DB = orm
}
