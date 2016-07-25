package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var odb orm.Ormer

func InitDB() {
	Connect()
	odb = orm.NewOrm()
	name := "default"
	force := false
	verbose := false
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Printf("init db is err ,this is %s ", err.Error())
	}
}
func Connect() {
	dns, _ := getConfig(1)
	fmt.Printf("数据库is %s", dns)
	err := orm.RegisterDataBase("default", "mysql", dns)
	if err != nil {
		fmt.Println("\n========>  数据库连接failed  <==========")
	} else {
		fmt.Println("\n========>  数据库连接sucess  <==========")
	}
}
func getConfig(flag int) (string, string) {
	var dns string
	db_host := beego.AppConfig.String("db_host")
	db_port := beego.AppConfig.String("db_port")
	db_user := beego.AppConfig.String("db_user")
	db_pass := beego.AppConfig.String("db_pass")
	db_name := beego.AppConfig.String("db_name")
	if flag == 1 {
		fmt.Println("========>     连接数据库     <==========")
		orm.RegisterDriver("mysql", orm.DRMySQL)
		dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Local", db_user, db_pass, db_host, db_port, db_name)
	} else {
		fmt.Println("========>     创建数据库     <==========")
		dns = fmt.Sprintf("%s:%s@tcp(%s:%s)?charset=utf8", db_user, db_pass, db_host, db_port)
	}

	return dns, db_name
}
