package models

import (
	// "fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type CompanyConfig struct {
	Id                  int
	Code                string `orm:"unique"` //公司代码
	DatabseInfo         string //连接数据库信息
	Addr                string //服务启动地址 默认0.0.0.0
	Httpport            int    //端口号
	DayMaxDownloadCount int    `orm:"default(150)"`           //每天下载资源的最大数量
	DayMaxHidDownload   int    `orm:"default(30)"`            //每天设备下载资源的最大数量
	ClearDataTimes      string `orm:"default(0 28 17 * * *)"` //执行自动清除任务的时间
	Status              int    `orm:"size(1);default(1)"`     //状态 1开启 0关闭
	ServerAddr          string //服务器地址
	Ormdebug            int    `orm:"size(1);default(0)"`    //状态 1开启 0关闭
	Testdebug           int    `orm:"size(1);default(0)"`    //状态 1开启 0关闭
	Maxuploadimages     int    `orm:"size(22);default(150)"` //最大上传图片个数(0表不受限制)
	Adminuser           int64  `orm:"default(51875511)"`     //超级管理员
	Recycle             int64  `orm:"default(999999)"`       //回收用户
}

func (this *CompanyConfig) TableName() string {
	return "company_config"
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(CompanyConfig))
}

//获取所有配置信息
func GetCompanyConfig(code string) (c CompanyConfig, err error) {
	o := orm.NewOrm()
	err = o.QueryTable("company_config").Filter("Code", code).Filter("Status", 1).One(&c)
	return c, err
}

// 获取公司信息
func GetAllCompanyConfig(code string) (c CompanyConfig) {
	o := orm.NewOrm()
	o.QueryTable("company_config").Filter("Code", code).One(&c)
	return c
}

// 更新公司状态
func UpdateCompanyStatus(code string, status int64) (num int64, err error) {
	model := orm.NewOrm()
	num, err = model.QueryTable("company_config").Filter("Code", code).Update(orm.Params{
		"Status": status,
	})
	return num, err
}

//根据Code值查询配置信息
func GetConfigInfo(code string) (companyconfig CompanyConfig) {
	o := orm.NewOrm()
	companyconfig = CompanyConfig{Code: code}
	err := o.Read(&companyconfig, "Code")
	if err == orm.ErrNoRows {
		beego.Error("can't find the config", err)
	} else if err == orm.ErrMissPK {
		beego.Error("can't find the PK", err)
	}
	return companyconfig
}

//增加配置信息
func AddCompanyconfigInfo(c *CompanyConfig) (int64, error) {
	o := orm.NewOrm()
	var companyconfig CompanyConfig
	companyconfig.Code = c.Code
	companyconfig.DatabseInfo = c.DatabseInfo
	companyconfig.Addr = c.Addr
	companyconfig.Httpport = c.Httpport
	companyconfig.DayMaxDownloadCount = c.DayMaxDownloadCount
	companyconfig.DayMaxHidDownload = c.DayMaxHidDownload
	companyconfig.ClearDataTimes = c.ClearDataTimes
	companyconfig.Status = c.Status
	companyconfig.ServerAddr = c.ServerAddr
	companyconfig.Ormdebug = c.Ormdebug
	companyconfig.Testdebug = c.Testdebug
	companyconfig.Maxuploadimages = c.Maxuploadimages
	companyconfig.Adminuser = c.Adminuser
	companyconfig.Recycle = c.Recycle
	id, err := o.Insert(companyconfig)
	if err != nil {
		beego.Error("insert is err", err)
	}
	return id, err
}
