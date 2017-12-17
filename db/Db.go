/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/16 1:25
 * @description: 
 */

package db

import (
	"os"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/gommon/log"
)

type dbconf struct {
	driver string
	dsn string
	network string
	maxconn int
	maxidle int
}

var reconnectNum int = 0

func GetConn() *gorm.DB {
	env := os.Getenv("APP_ENV")
	var conf dbconf = dbconf{}
	if env == "production" {
		conf = dbconf{
			driver: "mysql",
			dsn: "root:111111@tcp(Localhost:3306)/hahajok?timeout=90s&charset=utf8&loc=Local&parseTime=true",
			network: "tcp",
			maxconn: 500,
			maxidle: 300,
		}
	} else if env == "beta" {
		conf = dbconf{
			driver: "mysql",
			dsn: "root:111111@tcp(Localhost:3306)/hahajok?timeout=90s&charset=utf8&loc=Local&parseTime=true",
			network: "tcp",
			maxconn: 500,
			maxidle: 300,
		}
	} else {
		conf = dbconf{
			driver: "mysql",
			dsn: "root:111111@tcp(Localhost:3306)/hahajok?timeout=90s&charset=utf8&loc=Local&parseTime=true",
			network: "tcp",
			maxconn: 500,
			maxidle: 300,
		}
	}
	db, err := gorm.Open(conf.driver, conf.dsn)
	if err != nil || db.DB().Ping() != nil {
		if reconnectNum > 3 {
			log.Errorf("连接数据库三次失败:%v", err)
			panic(err)
		}
		reconnectNum += 1
		return GetConn()
	}
	reconnectNum = 0
	db.DB().SetMaxOpenConns(conf.maxconn)
	db.DB().SetMaxIdleConns(conf.maxidle)
	return db
}
