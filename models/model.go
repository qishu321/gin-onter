package models

import (
	"fmt"
	"gin-onter/conf"
	"gin-onter/models/models_lottery"
	"time"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB
var err error

func InitDb() {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.DbUser,
		conf.DbPassWord,
		conf.DbHost,
		conf.DbPort,
		conf.DbName,
	)
	Db, err = gorm.Open("mysql", dns)
	if err != nil {
		fmt.Println("数据库连接失败", err)
	}
	Db.SingularTable(true)
	Db.DB().SetMaxIdleConns(20)  //设置连接池，空闲
	Db.DB().SetMaxOpenConns(100) //打开
	Db.DB().SetConnMaxLifetime(time.Second * 30)
	Db.AutoMigrate(models_lottery.Lottery{},models_lottery.Prize{})
	Db.LogMode(true)

}
//func DB() (*gorm.DB) {
//	return Db
//}