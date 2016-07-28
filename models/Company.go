package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// "reflect"
)

//公司表
type Company struct {
	Id             int64
	Code           string `orm:"unique"` //公司代码
	Name           string //公司名称
	Status         int    `orm:"default(1)"`     //公司状态 0:关闭 1:正常
	ExpirationTime int64  `orm:"type(datetime)"` //公司过期时间
	AddTime        int64  `orm:"type(datetime)"` //公司添加时间
	LoginUrl       string //C 登录的所有地址

	UpdateUrl string //请求升级的URL

	SoftVersion   int64  //C 软件版本号
	SoftUpdateUrl string //C 软件更新
	SoftMd5       string //C 文件md5
	SoftDescribe  string //C 软件描述
	SoftCleardata int32  `orm:"default(0)"` //C 清空数据 0 清空，1不清空

	RomVersion   int64  //C ROM版本号
	RomUpdateUrl string //C ROM更新
	RomMd5       string // C文件md5
	RomDescribe  string // C rom描述
	RomCleardata int32  `orm:"default(0)"` //C 清空数据 0 清空，1不清空

	ApkVersion   int64  //JAVA 版本号
	ApkUpdateUrl string //JAVA 更新APK
	ApkMd5       string // Java 文件md5
	ApkDescribe  string // JAVA 升级描述
	ApkClerdata  int32  `orm:"default(0)"` //C 清空数据 0 清空，1不清空

	CustomerUrl   string //JAVA 下载联系人地址
	ImagesUrl     string //JAVA 图片地址
	WechatDataUrl string //JAVA 微信记录接口

	UserMaxCount  int64 //用户最大拥有设备数量
	TotalMaxCount int64 //公司总共用户设备最大数量

	FuncCustomer      int32 `orm:"default(1)"` //开启下载资源 (1 :默认开启 0:关闭)
	FuncImage         int32 `orm:"default(1)"` //开启下载图片 (1 :默认开启 0:关闭)
	FuncUploadWechat  int32 `orm:"default(1)"` //开启上传微信记录 (1 :默认开启 0:关闭)
	FuncLimitCustomer int32 `orm:"default(1)"` //开启资源限定(1:开启，0：关闭)

	FuncAutoAdd    int32             `orm:"default(1)"`      //开启自动添加用户 (1 :默认开启 0:关闭)
	FuncAutoSearch int32             `orm:"default(1)"`      //开启自动搜索用户 (1 :默认开启 0:关闭)
	FuncLocation   int32             `orm:"default(1)"`      //开启定位 (1 :默认开启 0:关闭)
	FuncShake      int32             `orm:"default(1)"`      //开启摇一摇 (1 :默认开启 0:关闭)
	CusttomerPush  int64             `orm:"default(0)"`      //开启推送资源公司(1:推送 0：默认不推送)
	CustomerUser   int64             `orm:"default(888888)"` //分发资源的帐号
	CustomerNum    int64             //分发资源量,资源限定
	MobileProvince []*MobileProvince `orm:"rel(m2m)"` //省份
}

func (this *Company) TableName() string {
	return "company"
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Company))
}

func GetCompany(code string) (*Company, bool) {

	o := orm.NewOrm()
	company := Company{Code: code}
	if err := o.Read(&company, "Code"); err != nil {
		return nil, false
	}

	return &company, true
}

/*func AddCompany() (int64, error) {
	omodel := orm.NewOrm()
	company := new(Company)
	company.Code = "001"
	company.Name = "武汉御诚"
	company.Status = 1
	company.ExpirationTime = 1433225246
	company.AddTime = 1433225246

	id, err := omodel.Insert(company)
	return id, err
}*/

//获取公司列表
func GetConpanyList(page int64, page_size int64, sort string) (companys []orm.Params, count int64) {
	omodel := orm.NewOrm()
	company := new(Company)
	qs := omodel.QueryTable(company)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort).Values(&companys)
	count, _ = qs.Count()
	return companys, count
}

//删除公司信息
func DelCompanyById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Company{Id: Id})
	return status, err
}

//更新公司信息
func UpdateCompany(u *Company) (int64, error) {
	omodel := orm.NewOrm()

	company := make(orm.Params)
	company["Code"] = u.Code
	company["Name"] = u.Name
	company["Status"] = u.Status
	company["ExpirationTime"] = u.ExpirationTime
	company["UserMaxCount"] = u.UserMaxCount
	company["TotalMaxCount"] = u.TotalMaxCount
	company["AddTime"] = u.AddTime
	company["SoftCleardata"] = u.SoftCleardata
	company["ApkClerdata"] = u.ApkClerdata
	company["RomCleardata"] = u.RomCleardata
	company["LoginUrl"] = u.LoginUrl
	company["SoftVersion"] = u.SoftVersion
	company["SoftUpdateUrl"] = u.SoftUpdateUrl
	company["SoftMd5"] = u.SoftMd5
	company["SoftDescribe"] = u.SoftDescribe
	company["RomVersion"] = u.RomVersion
	company["RomUpdateUrl"] = u.RomUpdateUrl
	company["RomMd5"] = u.RomMd5
	company["RomDescribe"] = u.RomDescribe
	company["ApkVersion"] = u.ApkVersion
	company["ApkUpdateUrl"] = u.ApkUpdateUrl
	company["ApkMd5"] = u.ApkMd5
	company["ApkDescribe"] = u.ApkDescribe
	company["ImagesUrl"] = u.ImagesUrl
	company["UpdateUrl"] = u.UpdateUrl
	company["CustomerUrl"] = u.CustomerUrl
	company["WechatDataUrl"] = u.WechatDataUrl
	company["FuncCustomer"] = u.FuncCustomer
	company["FuncImage"] = u.FuncImage
	company["FuncUploadWechat"] = u.FuncUploadWechat
	company["FuncLimitCustomer"] = u.FuncLimitCustomer
	company["FuncAutoAdd"] = u.FuncAutoAdd
	company["FuncAutoSearch"] = u.FuncAutoSearch
	company["FuncLocation"] = u.FuncLocation
	company["FuncShake"] = u.FuncShake

	company["CusttomerPush"] = u.CusttomerPush
	if u.CusttomerPush == 1 {
		company["CustomerUser"] = u.CustomerUser
	}

	var table Company
	num, err := omodel.QueryTable(table).Filter("Id", u.Id).Update(company)
	return num, err
}

//获取公司信息
func GetCompanyInfoById(id int64) (company Company) {
	o := orm.NewOrm()
	company = Company{Id: id}
	err := o.Read(&company)
	if err != nil {
		fmt.Println(nil)
	}
	return company
}

//通过code值查询公司
func GetCompanyInfoByCode(code string) (company Company) {
	o := orm.NewOrm()
	company = Company{Code: code}
	err := o.Read(&company, "Code")
	if err == orm.ErrNoRows {
		beego.Error("can't find the company", err)
	} else if err == orm.ErrMissPK {
		beego.Error("can't find the PK", err)
	}
	return company
}

//添加公司信息
func AddCompany(u *Company) (int64, error) {
	o := orm.NewOrm()

	var company Company
	company.Code = u.Code
	company.Name = u.Name
	company.Status = u.Status
	company.ExpirationTime = u.ExpirationTime
	company.UserMaxCount = u.UserMaxCount
	company.TotalMaxCount = u.TotalMaxCount
	company.AddTime = u.AddTime
	company.SoftCleardata = u.SoftCleardata
	company.ApkClerdata = u.ApkClerdata
	company.RomCleardata = u.RomCleardata
	company.LoginUrl = u.LoginUrl
	company.SoftVersion = u.SoftVersion
	company.SoftUpdateUrl = u.SoftUpdateUrl
	company.SoftMd5 = u.SoftMd5
	company.SoftDescribe = u.SoftDescribe
	company.RomVersion = u.RomVersion
	company.RomUpdateUrl = u.RomUpdateUrl
	company.RomMd5 = u.RomMd5
	company.RomDescribe = u.RomDescribe
	company.ApkVersion = u.ApkVersion
	company.ApkUpdateUrl = u.ApkUpdateUrl
	company.ApkMd5 = u.ApkMd5
	company.ApkDescribe = u.ApkDescribe
	company.ImagesUrl = u.ImagesUrl
	company.UpdateUrl = u.UpdateUrl
	company.CustomerUrl = u.CustomerUrl
	company.WechatDataUrl = u.WechatDataUrl
	company.FuncCustomer = u.FuncCustomer
	company.FuncImage = u.FuncImage
	company.FuncUploadWechat = u.FuncUploadWechat
	company.FuncLimitCustomer = u.FuncLimitCustomer

	company.FuncAutoAdd = u.FuncAutoAdd
	company.FuncAutoSearch = u.FuncAutoSearch
	company.FuncLocation = u.FuncLocation
	company.FuncShake = u.FuncShake
	company.CusttomerPush = u.CusttomerPush
	if u.CusttomerPush == 1 {
		company.CustomerUser = u.CustomerUser
	}

	num, err := o.Insert(&company)
	return num, err
}

// 获取所有公司
func GetAllCompany() (companys []orm.Params) {
	omodel := orm.NewOrm()
	company := new(Company)
	qs := omodel.QueryTable(company)
	qs.OrderBy("Id").Values(&companys)
	return companys
}

// 停止公司服务
func StopCompany(code string, status int64) (num int64, err error) {
	model := orm.NewOrm()
	num, err = model.QueryTable("company").Filter("Code", code).Update(orm.Params{
		"Status": int(status),
	})
	return num, err
}

// 查询条件公司数量
func GetFieldCompanyNum(status int) ([]orm.Params, int64, error) {
	var companylist []orm.Params
	model := orm.NewOrm()
	company := new(Company)
	qs := model.QueryTable(company)

	var cond *orm.Condition
	cond = orm.NewCondition()

	if status > 0 {
		cond = cond.And("Status", status)
	}
	qs.SetCond(cond).Values(&companylist)
	count, err := qs.SetCond(cond).Count()
	return companylist, count, err
}

// 更新公司和省份资源
func UpdateCompanyProvice(companyId int64, proviceId int64) (int64, error) {
	o := orm.NewOrm()
	company := Company{Id: companyId}
	provice := MobileProvince{Id: proviceId}
	m2m := o.QueryM2M(&company, "MobileProvince")
	num, err := m2m.Add(&provice)
	// m2m.
	return num, err
}

// 删除公司中的provice
func DelCompanyProvice(companyId int64) (int64, error) {
	o := orm.NewOrm()
	num, err := o.QueryTable("company_mobile_provinces").Filter("company_id", companyId).Delete()
	return num, err
}

// 更新公司信息
func UpdateCompanyInfo(id int64, companyname string, customerpush int64, customeruser int64, CustomerNum int64) (int64, error) {
	o := orm.NewOrm()
	num, err := o.QueryTable("company").Filter("Id", id).Update(orm.Params{
		"Name":          companyname,
		"CusttomerPush": customerpush,
		"CustomerUser":  customeruser,
		"CustomerNum":   CustomerNum,
	})
	return num, err
}

// 公司ID查询出省份
func ReloadCompanyProvice(id int64) (provice []orm.Params, count int64) {
	o := orm.NewOrm()
	count, _ = o.QueryTable("mobile_province").Filter("Company__Company__Id", id).Values(&provice)
	return provice, count

}

//根据Id值更改公司状态
func ChangeStatusById(id int64, status int) (int64, error) {
	o := orm.NewOrm()
	num, err := o.QueryTable("company").Filter("Id", id).Update(orm.Params{
		"Status": status,
	})
	return num, err
}

//根据Id值查询公司,返回一个数组
func GetCompanyById(Id int64) (company []orm.Params) {
	o := orm.NewOrm()
	_, err := o.QueryTable("company").Filter("Id", Id).Values(&company)
	if err != nil {
		beego.Error("can't find the company", err)
	}
	return company
}

//根据公司名字进行查询
func GetCompanyByCompanyName(companyName string) (company Company) {
	o := orm.NewOrm()
	company = Company{Name: companyName}
	err := o.Read(&company, "Name")
	if err == orm.ErrNoRows {
		beego.Error("can't find the company", err)
	} else if err == orm.ErrMissPK {
		beego.Error("can't find the PK", err)
	}
	return company
}

//company_mobile_provinces
//根据省份和公司代码查询公司信息
func GetCompanyByProvinceAndCode(provinceName string, code string) (company Company) {
	provinceid := GetCompanyIdByCompanyName(provinceName) //根据省份名称取得省份id，
	companyids := GetcompanyIdByProvinceId(provinceid)    //根据省份id取得公司id
	var companys []orm.Params
	for i := 0; i < len(companyids); i++ { //遍历循环取出所有的公司然后用code值排除不符合条件的公司。随后展现到前台
		companyid := companyids[i]
		companys = GetCompanyById(companyid)
		for _, v := range companys {
			companyCodes := v["Code"].(string)
			if companyCodes == code {
				company = GetCompanyInfoByCode(companyCodes)
			}
		}
	}
	return company
}

//根据省份和公司名查询公司信息
func GetCompanyByProvinceAndName(provinceName string, name string) (company Company) {
	provinceid := GetCompanyIdByCompanyName(provinceName) //根据省份名称取得省份id，
	companyids := GetcompanyIdByProvinceId(provinceid)    //根据省份id取得公司id
	var companys []orm.Params
	for i := 0; i < len(companyids); i++ {
		companyid := companyids[i]
		companys = GetCompanyById(companyid)
		for _, v := range companys {
			companyNames := v["Name"].(string)
			if companyNames == name {
				companyCode := v["Code"].(string)
				company = GetCompanyInfoByCode(companyCode)
			}
		}

	}
	return company
}

//根据省份和状态查询公司信息
func GetCompanyByProvinceAndStatus(provinceName string, statu int) (company []Company) {
	provinceid := GetCompanyIdByCompanyName(provinceName) //根据省份名称取得省份id，
	companyids := GetcompanyIdByProvinceId(provinceid)    //根据省份id取得公司id
	var companys []orm.Params
	for i := 0; i < len(companyids); i++ { //遍历循环取出所有的公司然后用code值排除不符合条件的公司。随后展现到前台
		companyid := companyids[i]
		companys = GetCompanyById(companyid)
		for _, v := range companys {
			sta := v["Status"].(int64)
			status := int64(statu)
			if sta == status {
				companyCode := v["Code"].(string)
				companySingle := GetCompanyInfoByCode(companyCode)
				company = append(company, companySingle)
			}
		}
	}
	return company
}

//根据公司状态查询公司信息
func GetCompanyByStatus(status int) (company []orm.Params, num int64) {
	o := orm.NewOrm()
	var companys []orm.Params
	nums, err := o.QueryTable("company").Filter("Status", status).Values(&companys)
	if err != nil {
		beego.Error("can't find the company", err)
	}
	return companys, nums

}

//查询所有公司的数量
func GetAllCompanyNum() (int64, error) {
	o := orm.NewOrm()
	count, err := o.QueryTable("company").Count()
	return count, err
}

//根据id值更新公司属性
func UpdateCompanysById(Id int64, CompanyName string, CusttomerPush int64, CustomerUser int64, CustomerNum int64) (int64, error) {
	beego.Debug(Id, CompanyName, CusttomerPush, CustomerUser, CustomerNum)
	o := orm.NewOrm()
	num, err := o.QueryTable("company").Filter("Id", Id).Update(orm.Params{
		"Name":          CompanyName,
		"CusttomerPush": CusttomerPush,
		"CustomerUser":  CustomerUser,
		"CustomerNum":   CustomerNum,
	})
	return num, err
}
