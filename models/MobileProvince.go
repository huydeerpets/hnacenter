package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//客户资源批次记录表
type MobileProvince struct {
	Id           int64
	Provinceid   int64      `orm:"index"` //省份
	Provincename string     //省份名称
	Company      []*Company `orm:"reverse(many)"` //
}

func (this *MobileProvince) TableName() string {
	return "mobile_province"
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(MobileProvince))
}

//根据省份id取得省份信息
func GetMobileProvice(provinceid int64) (*MobileProvince, bool) {
	o := orm.NewOrm()
	province := MobileProvince{Provinceid: provinceid}
	if err := o.Read(&province, "Provinceid"); err != nil {
		return nil, false
	}
	return &province, true
}

//取得所有的不重复省份
func GetAllProvice() []MobileProvince {
	o := orm.NewOrm()
	var province []MobileProvince
	_, err := o.QueryTable("mobile_province").All(&province)
	if err != nil {
		beego.Error("can't find the province", err)
	}
	return province
}

//根据省份名称取得省份id，
func GetCompanyIdByCompanyName(name string) (provinceid int64) {
	o := orm.NewOrm()
	var mobilprovince []orm.Params
	_, err := o.QueryTable("mobile_province").Filter("Provincename", name).Values(&mobilprovince)
	if err != nil {
		beego.Error("province is abnormal", err)
	} else {
		for _, v := range mobilprovince {
			provinceid = v["Provinceid"].(int64)
		}
	}
	return provinceid
}

//根据省份id取得公司id
func GetcompanyIdByProvinceId(provinceid int64) (companyid []int64) {
	o := orm.NewOrm()
	var companyAndProvince []orm.Params
	var companyids []int64
	_, err := o.QueryTable("company_mobile_provinces").Filter("mobile_province_id", provinceid).Values(&companyAndProvince)
	if err != nil {
		beego.Error("no company on this province", err)
	} else {
		for i := 0; i < len(companyAndProvince); i++ {
			companyAndProvinces := companyAndProvince[i]
			companyids = append(companyids, companyAndProvinces["Company"].(int64))
		}
	}
	return companyids
}
