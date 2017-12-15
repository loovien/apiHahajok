/**
 * @website: https://vvotm.github.io
 * @author luowen<bigpao.luo@gmail.com>
 * @date 2017/12/16 1:25
 * @description: 
 */

package db

import (
	"os"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type dbconf struct {
	driver string
	dsn string
	network string
	maxconn int
	maxidle int
}

var reconnectNum int = 0

func GetConn()  *sql.DB {
	env := os.Getenv("APP_ENV")
	var conf dbconf = dbconf{}
	if env == "production" {
		conf = dbconf{
			driver: "mysql",
			dsn: "root:111111@tcp(localhost:3306)/hahajok?timeout=90s&charset=utf8",
			network: "tcp",
			maxconn: 500,
			maxidle: 300,
		}
	} else if env == "beta" {
		conf = dbconf{
			driver: "mysql",
			dsn: "root:111111@tcp(localhost:3306)/hahajok?timeout=90s&charset=utf8",
			network: "tcp",
			maxconn: 500,
			maxidle: 300,
		}
	} else {
		conf = dbconf{
			driver: "mysql",
			dsn: "root:111111@tcp(localhost:3306)/hahajok?timeout=90s&charset=utf8",
			network: "tcp",
			maxconn: 500,
			maxidle: 300,
		}
	}
	db, err := sql.Open(conf.driver, conf.dsn)
	if err != nil || db.Ping() != nil {
		if reconnectNum > 3 {
			panic(err)
		}
		reconnectNum += 1
		return GetConn()
	}
	reconnectNum = 0
	db.SetMaxOpenConns(conf.maxconn)
	db.SetMaxIdleConns(conf.maxidle)
	return db
}
