package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	utils "hnacenter/src/utils"
	"time"
)

//客户资源批次记录表
type CustomerDistribute struct {
	Id            int64
	AddTime       string        //添加时间
	AddUser       int64         //添加者
	DistributeNum int64         `orm:"default(0)"`         //批次总数
	SuccessNum    int64         `orm:"default(0)"`         //成功上传数量
	FailNum       int64         `orm:"default(0)"`         //失败数量
	RepeatNum     int64         `orm:"default(0)"`         //重复数量
	TotalNum      int64         `orm:"default(0)"`         //数据库总数量
	Type          *CustomerType `orm:"rel(fk);default(1)"` //资源类型
}

func (this *CustomerDistribute) TableName() string {
	return "customer_distribute"
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(CustomerDistribute))
}

// 批次查看
func ListCustomerDis() (listdis []orm.Params) {
	model := orm.NewOrm()
	_, err := model.QueryTable("customer_distribute").OrderBy("-id").Values(&listdis)
	if err != nil {
		beego.Error(err)
	}
	return listdis
}

// 增加导入批次
func AddCustomerDis(uname int64, _type int64) (int64, error) {
	modal := orm.NewOrm()
	var customerdis CustomerDistribute
	nowtime := utils.Format(time.Now(), "yyyy-MM-dd HH:mm:ss")
	customerdis.AddTime = nowtime
	customerdis.AddUser = uname
	customerdis.Type = &CustomerType{Id: _type}

	beego.Info(customerdis)

	id, err := modal.Insert(&customerdis)
	if err != nil {
		beego.Error(err)
		return 0, err
	}
	return id, nil
}

// 更新批次信息
func UpdateDistribute(d CustomerDistribute, fields ...string) int64 {
	o := orm.NewOrm()
	num, _ := o.Update(&d, fields...)
	return num
}

// 查询左后一个批次
func GetLastDistribute() (dis CustomerDistribute, err error) {
	model := orm.NewOrm()
	err = model.QueryTable("customer_distribute").OrderBy("-Id").Limit(1).One(&dis)
	return dis, err
}
