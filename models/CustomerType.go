package models

import (
	// "fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type CustomerType struct {
	Id         int64
	Name       string
	CreateTime time.Time `orm:"type(datetime);auto_now_add"`
}

func (this *CustomerType) TableName() string {
	return "customer_type"
}
func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(CustomerType))
}

func AddCustomerType(r CustomerType) int64 {
	o := orm.NewOrm()
	num, err := o.Insert(&r)
	if err != nil {
		beego.Error(err)
		return 0
	}
	return num
}

func EditCustomerType(id int64, name string) int64 {
	o := orm.NewOrm()
	r := make(orm.Params)
	r["Name"] = name
	num, err := o.QueryTable("customer_type").Filter("Id", id).Update(r)
	if err != nil {
		beego.Error(err)
		return 0
	}
	return num
}

func CustomerTypeList() (list []orm.Params) {
	o := orm.NewOrm()
	o.QueryTable("customer_type").Values(&list)
	return list
}

func GetCusTypeById(id int64) (c CustomerType) {
	o := orm.NewOrm()
	o.QueryTable("customer_type").Filter("Id", id).One(&c)
	return c
}

func GetCountTypeList() (int64, error) {
	o := orm.NewOrm()
	count, err := o.QueryTable("customer_type").Count()
	return count, err
}
