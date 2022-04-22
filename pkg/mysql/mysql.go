package mysql

import (
	"fmt"
	"github.com/FFFcomewhere/University-course-selection/test-demo/model"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type config struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
}

var engine *xorm.Engine

//初始化表单
func InitDBTables(tableName string) {
	//TODO 后面完善,根据表单名初始化
	err := engine.Sync2(new(model.User))
	if err != nil {
		fmt.Println(err)
	}
}

//初始化数据库
func InitMysql(host, port, username, password, dbname string) {
	//载入配置信息
	dataSourceName := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"

	//fmt.Println(dataSourceName)
	var err error
	engine, err = xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		log.Println(err)
		return
	}
}

func Engine() *xorm.Engine {
	return engine
}
