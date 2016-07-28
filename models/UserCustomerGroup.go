package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//公司组别表
type UserCustomerGroup struct {
	Id        int64
	Name      string
	UserGroup string `orm:"type(text)"`  //组成员
	UserCount int    `orm:"default(0);"` //成员数量
}

func (this *UserCustomerGroup) TableName() string {
	return "user_customer_group"
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(UserCustomerGroup))
}

// 资源组列表
func GetuCustomerGroupList(page int64, page_size int64, id string) (ucustomergroup []orm.Params, count int64) {
	omodel := orm.NewOrm()
	uCustomer := new(UserCustomerGroup)
	qs := omodel.QueryTable(uCustomer)

	qs.Limit(page_size, page).OrderBy(id).Values(&ucustomergroup)

	count, _ = qs.Count()
	return ucustomergroup, count
}

// 所有资源组列表
func GetAlluCustomerGroupList() (ucustomergroup []orm.Params) {
	o := orm.NewOrm()
	uCustomer := new(UserCustomerGroup)
	o.QueryTable(uCustomer).Values(&ucustomergroup)
	return ucustomergroup
}

// 增加资源组
func AdduCustomerGroup(ucustomergroup *UserCustomerGroup) (int64, error) {
	omodel := orm.NewOrm()
	inserid, err := omodel.Insert(ucustomergroup)
	return inserid, err
}

// 删除资源组
func DeluCustomerGroup(id int64) (int64, error) {
	omodel := orm.NewOrm()
	delid, err := omodel.Delete(&UserCustomerGroup{Id: id})
	return delid, err
}

// 根据Id查询组内容
func ReaduCustomerGroup(id int64) (ucustomergroup UserCustomerGroup, err error) {
	omodel := orm.NewOrm()
	err = omodel.QueryTable("user_customer_group").Filter("Id", id).One(&ucustomergroup)
	return ucustomergroup, err
}

// 修改资源组
func UpdateuCustomerGroup(ucustomergroup *UserCustomerGroup, ucustomerid int64) (int64, error) {
	omodel := orm.NewOrm()
	updateid, err := omodel.Update(ucustomergroup)
	return updateid, err
}

//获取指定id数组的对象
func GetCustomerGroupByIdArr(ids []string) (u []UserCustomerGroup) {
	o := orm.NewOrm()
	o.QueryTable("user_customer_group").Filter("Id__in", ids).All(&u)
	return u
}

//根据name值查找资源组
func GetUserGroupByName(name string) (usergroup UserCustomerGroup) {
	o := orm.NewOrm()
	usergroup = UserCustomerGroup{Name: name}
	err := o.Read(&usergroup, "Name")
	if err != nil {
		beego.Error("can't find the usergroup", err)
	}
	return usergroup
}
