package test_demo

import (
	"fmt"
	"github.com/FFFcomewhere/University-course-selection/test-demo/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"testing"
)

func TestMysql(t *testing.T) {
	dbSource := "root:password@tcp(127.0.0.1:3306)/user?charset=utf8"

	xormEngine, err := xorm.NewEngine("mysql", dbSource)
	if err != nil {
		fmt.Println(err)
	}
	user := &model.User{
		Id: "2",
	}
	status, err := xormEngine.Get(user)

	fmt.Println(status)

	//DB, _ := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/user?charset=utf8")
	//
	////设置数据库最大连接数
	//DB.SetConnMaxLifetime(100)
	////设置上数据库最大闲置连接数
	//DB.SetMaxIdleConns(10)
	////验证连接
	//if err := DB.Ping(); err != nil {
	//	fmt.Println("open database fail")
	//	return
	//}
	//fmt.Println("connnect success")
}
