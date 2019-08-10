package models

import (
	"fmt"
	"gin-server/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

type GormLogger struct{}

func (*GormLogger) Print(v ...interface{}) {
	if v[0] == "sql" {
		log.WithFields(log.Fields{"module": "gorm", "type": "sql"}).Print(v[3], v[4])
	}
	if v[0] == "log" {
		log.WithFields(log.Fields{"module": "gorm", "type": "log"}).Print(v[2])
	}
}

var MysqlDb *gorm.DB

// 初始化数据库连接
func init() {
	mysqlConf := config.Config.MysqlConfig
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", mysqlConf.User,
		mysqlConf.Password, mysqlConf.Host, mysqlConf.Port, mysqlConf.Url)
	//dbUrl := fmt.Sprintf("%s:%s@/%s", mysqlConf.User,
	//	mysqlConf.Password, mysqlConf.Url)
	db, err := gorm.Open("mysql", dbUrl)
	if err != nil {
		log.Fatalf("%v, 数据库链接失败: %s！", dbUrl, err)
	}

	log.Info("数据库链接成功！")
	MysqlDb = db
	MysqlDb.SetLogger(&GormLogger{})
	MysqlDb.LogMode(true)
}

func CloseDB() {
	if MysqlDb != nil {
		MysqlDb.Close()
		log.Infof("关闭数据库链接!")
	}
}
