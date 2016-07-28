package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//客户资源批次记录表
type Mobile struct {
	Id               int64
	Mobilenumber     int64  //判断前面几位
	Mobilearea       string //中文 地名
	Mobileareanumber string //手机卡类
	Mobilenet        int64  //区号
	Mobilepost       int64  //postcode
	Province         int64  //省份
}

func (this *Mobile) TableName() string {
	return "mobile"
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Mobile))
}

func CheckMobileArea(phone string) Mobile {
	model := orm.NewOrm()
	subphone := phone[0:7]
	var mobile Mobile
	err := model.QueryTable("mobile").Filter("Mobilenumber", subphone).One(&mobile)
	if err != nil {
		// fmt.Println(err)
		beego.Error(err)
	}
	// fmt.Println(mobile)
	return mobile
}

// 随机获取手机号
func GetRandMobileNumber(phone string) (mobileinfo []orm.Params, err error) {
	model := orm.NewOrm()
	_, err = model.QueryTable("mobile").Filter("mobileareanumber__contains", "移动").Filter("mobilenumber__startswith", phone).Values(&mobileinfo)
	return mobileinfo, err
}
