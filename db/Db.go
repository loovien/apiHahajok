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
	"sync"
)

type dbconf struct {
	driver string
	dsn string
	network string
	maxconn int
	maxidle int
}

type DbMgr struct {
	dbConf dbconf
	dbConn *gorm.DB
	lock *sync.RWMutex
}

var (
	dbmgr *DbMgr = nil
	dbOnce *sync.Once = &sync.Once{}
)

func NewDbMgr() *DbMgr {
	if dbmgr == nil {
		dbOnce.Do(func() {
			dbmgr = &DbMgr{
				lock: &sync.RWMutex{},
			}
		})
	}
	return dbmgr
}
func (d *DbMgr) Init()  {
	defer d.lock.Unlock()
	d.lock.Lock()
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
	conn, err := gorm.Open(conf.driver, conf.dsn)
	if err != nil {
		log.Fatal(err)
	}
	conn.DB().SetMaxOpenConns(conf.maxconn)
	conn.DB().SetMaxIdleConns(conf.maxidle)
	d.dbConn = conn
	d.dbConf = conf
}

func GetConn() *gorm.DB {
	defer dbmgr.lock.Unlock()
	dbmgr.lock.Lock()
	if err := dbmgr.dbConn.DB().Ping(); err != nil {
		conn, err := gorm.Open(dbmgr.dbConf.driver, dbmgr.dbConf.dsn)
		if err != nil {
			log.Fatal(err)
		}
		conn.DB().SetMaxOpenConns(dbmgr.dbConf.maxconn)
		conn.DB().SetMaxIdleConns(dbmgr.dbConf.maxidle)
		dbmgr.dbConn = conn
	}
	return dbmgr.dbConn
}
