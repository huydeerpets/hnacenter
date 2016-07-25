package controllers

import (
	"github.com/astaxie/beego"
	m "hnacenter/models"
	"time"
)

type CompanyController struct {
	LoginController
}

//取得公司列表
func (this *CompanyController) GetCompanyList() {
	provices := m.GetAllProvice()
	this.Data["provice"] = provices
	if this.IsAjax() {
		sEcho := this.GetString("sEcho")
		iStart, _ := this.GetInt64("iDisplayStart")
		iLength, _ := this.GetInt64("iDisplayLength")
		companys, count := m.GetConpanyList(iStart, iLength, "Id")

		for _, company := range companys {
			switch company["Status"] {
			case int64(0):
				company["Statusname"] = "关闭"
			case int64(1):
				company["Statusname"] = "正常"
			default:
				company["Statusname"] = "异常"
			}
		}

		data := make(map[string]interface{})
		data["aaData"] = companys
		data["iTotalDisplayRecords"] = count
		data["iTotalRecords"] = iLength
		data["sEcho"] = sEcho
		this.Data["json"] = &data
		this.ServeJSON()
	} else {
		this.Data["ActionUrl"] = beego.AppConfig.String("ActionUrl")
		this.TplName = "company/index.html"
	}

}

//添加公司
func (this *CompanyController) AddCompany() {
	isAction := this.GetString("isAction")
	if isAction == "0" {
		Code := this.GetString("Code")                     //公司代码
		Name := this.GetString("Name")                     //公司名称
		Status, _ := this.GetInt("Status")                 //状态
		Addtime := this.GetString("AddTime")               //添加时间
		Expiration := this.GetString("ExpirationTime")     //过期时间
		Usermaxcount, _ := this.GetInt64("UserMaxCount")   //个人设备最大数量
		Totalmaxcount, _ := this.GetInt64("TotalMaxCount") //公司设备最大数量
		Loginurl := this.GetString("LoginUrl")             //设备登陆地址
		Customerurl := this.GetString("CustomerUrl")       //下载联系人地址
		Wechatdataurl := this.GetString("WechatDataUrl")   //微信记录接口
		Imagesurl := this.GetString("ImagesUrl")           //图片地址
		Updateurl := this.GetString("UpdateUrl")           //软件升级地址

		Softversion, _ := this.GetInt64("SoftVersion")     //软件版本号
		Softupdateurl := this.GetString("SoftUpdateUrl")   //软件更新地址
		Softmd5 := this.GetString("SoftMd5")               //文件MD5值
		Softdescribe := this.GetString("SoftDescribe")     //软件描述
		Softcleardata, _ := this.GetInt32("SoftCleardata") //软件清空数据

		Romversion, _ := this.GetInt64("RomVersion")     //Rom版本
		Romupdateurl := this.GetString("RomUpdteUrl")    //Rom更新地址
		Rommd5 := this.GetString("RomMd5")               //rom文件MD5值
		Romdescribe := this.GetString("RomDescribe")     //rom描述
		Romcleardata, _ := this.GetInt32("RomCleardata") //Rom清空数据

		Apkversion, _ := this.GetInt64("ApkVersion")    //APK版本号
		Apkupdateurl := this.GetString("ApkUpdateUrl")  //Apk更新地址
		Apkmd5 := this.GetString("ApkMd5")              //apk文件MD5值
		Apkdescribe := this.GetString("ApkDescribe")    //apk描述
		Apkcleardata, _ := this.GetInt32("ApkClerdata") //apk清空数据
		//1开启  0关闭
		FunCustomer, _ := this.GetInt32("FuncCustomer")            //下载资源开关
		FuncImage, _ := this.GetInt32("FuncImage")                 //下载图片开关
		FuncUploadWechat, _ := this.GetInt32("FuncUploadWechat")   //上传微信记录开关
		FuncLimitCustomer, _ := this.GetInt32("FuncLimitCustomer") //资源限定开关
		FuncAutoAdd, _ := this.GetInt32("FuncAutoAdd")             //自动添加用户开关
		FuncAutoSearch, _ := this.GetInt32("FuncAutoSearch")       //自动搜索用户开关
		FuncLocation, _ := this.GetInt32("FuncLocation")           //定位开关
		FuncShake, _ := this.GetInt32("FuncShake")                 //摇一摇开关
		CusttomerPush, _ := this.GetInt64("CusttomerPush")         //资源推送开关
		Expir, _ := time.Parse("2006-01-02 15:04:05", Expiration)
		addtime, _ := time.Parse("2006-01-02 15:04:05", Addtime)

		company := m.GetCompanyInfoByCode(Code)
		if company.Id == 0 {
			company.Code = Code
			company.Name = Name
			company.Status = Status
			company.ExpirationTime = Expir.Unix()
			company.AddTime = addtime.Unix()
			company.LoginUrl = Loginurl
			company.UpdateUrl = Updateurl
			company.SoftVersion = Softversion
			company.SoftUpdateUrl = Softupdateurl
			company.SoftMd5 = Softmd5
			company.SoftDescribe = Softdescribe
			company.SoftCleardata = Softcleardata
			company.RomVersion = Romversion
			company.RomUpdateUrl = Romupdateurl
			company.RomMd5 = Rommd5
			company.RomDescribe = Romdescribe
			company.RomCleardata = Romcleardata
			company.ApkVersion = Apkversion
			company.ApkUpdateUrl = Apkupdateurl
			company.ApkMd5 = Apkmd5
			company.ApkDescribe = Apkdescribe
			company.ApkClerdata = Apkcleardata
			company.CustomerUrl = Customerurl
			company.ImagesUrl = Imagesurl
			company.WechatDataUrl = Wechatdataurl
			company.UserMaxCount = Usermaxcount
			company.TotalMaxCount = Totalmaxcount
			company.FuncCustomer = FunCustomer
			company.FuncImage = FuncImage
			company.FuncUploadWechat = FuncUploadWechat
			company.FuncLimitCustomer = FuncLimitCustomer
			company.FuncAutoAdd = FuncAutoAdd
			company.FuncAutoSearch = FuncAutoSearch
			company.FuncLocation = FuncLocation
			company.FuncShake = FuncShake
			company.CusttomerPush = CusttomerPush
			_, err := m.AddCompany(&company)
			if err != nil {
				beego.Error("CompanyController insert is err", err)
				return
			}
		} else {
			this.Alert("添加成功", "../company/list")
		}
		this.Ctx.Redirect(302, "/hnacenter/company/add")
	} else {
		this.TplName = "company/add.html"
	}
}

//修改公司信息
func (this *CompanyController) EditCompany() {
	isAction := this.GetString("isAction")
	Id, _ := this.GetInt64("Id")
	company := m.GetCompanyInfoById(Id)
	if isAction == "0" {
		Code := this.GetString("Code")                     //公司代码
		Name := this.GetString("Name")                     //公司名称
		Status, _ := this.GetInt("Status")                 //状态
		Addtime := this.GetString("AddTime")               //添加时间
		Expiration := this.GetString("ExpirationTime")     //过期时间
		Usermaxcount, _ := this.GetInt64("UserMaxCount")   //个人设备最大数量
		Totalmaxcount, _ := this.GetInt64("TotalMaxCount") //公司设备最大数量
		Loginurl := this.GetString("LoginUrl")             //设备登陆地址
		Customerurl := this.GetString("CustomerUrl")       //下载联系人地址
		Wechatdataurl := this.GetString("WechatDataUrl")   //微信记录接口
		Imagesurl := this.GetString("ImagesUrl")           //图片地址
		Updateurl := this.GetString("UpdateUrl")           //软件升级地址

		Softversion, _ := this.GetInt64("SoftVersion")     //软件版本号
		Softupdateurl := this.GetString("SoftUpdateUrl")   //软件更新地址
		Softmd5 := this.GetString("SoftMd5")               //文件MD5值
		Softdescribe := this.GetString("SoftDescribe")     //软件描述
		Softcleardata, _ := this.GetInt32("SoftCleardata") //软件清空数据

		Romversion, _ := this.GetInt64("RomVersion")     //Rom版本
		Romupdateurl := this.GetString("RomUpdteUrl")    //Rom更新地址
		Rommd5 := this.GetString("RomMd5")               //rom文件MD5值
		Romdescribe := this.GetString("RomDescribe")     //rom描述
		Romcleardata, _ := this.GetInt32("RomCleardata") //Rom清空数据

		Apkversion, _ := this.GetInt64("ApkVersion")    //APK版本号
		Apkupdateurl := this.GetString("ApkUpdateUrl")  //Apk更新地址
		Apkmd5 := this.GetString("ApkMd5")              //apk文件MD5值
		Apkdescribe := this.GetString("ApkDescribe")    //apk描述
		Apkcleardata, _ := this.GetInt32("ApkClerdata") //apk清空数据
		//1开启  0关闭
		FunCustomer, _ := this.GetInt32("FuncCustomer")            //下载资源开关
		FuncImage, _ := this.GetInt32("FuncImage")                 //下载图片开关
		FuncUploadWechat, _ := this.GetInt32("FuncUploadWechat")   //上传微信记录开关
		FuncLimitCustomer, _ := this.GetInt32("FuncLimitCustomer") //资源限定开关
		FuncAutoAdd, _ := this.GetInt32("FuncAutoAdd")             //自动添加用户开关
		FuncAutoSearch, _ := this.GetInt32("FuncAutoSearch")       //自动搜索用户开关
		FuncLocation, _ := this.GetInt32("FuncLocation")           //定位开关
		FuncShake, _ := this.GetInt32("FuncShake")                 //摇一摇开关
		CusttomerPush, _ := this.GetInt64("CusttomerPush")         //资源推送开关
		CustomerUser, _ := this.GetInt64("CustomerUser")           //资源管理账号
		Expir, _ := time.Parse("2006-01-02 15:04:05", Expiration)
		addtime, _ := time.Parse("2006-01-02 15:04:05", Addtime)

		company.Code = Code
		company.Name = Name
		company.Status = Status
		company.ExpirationTime = Expir.Unix()
		company.AddTime = addtime.Unix()
		company.LoginUrl = Loginurl
		company.UpdateUrl = Updateurl
		company.SoftVersion = Softversion
		company.SoftUpdateUrl = Softupdateurl
		company.SoftMd5 = Softmd5
		company.SoftDescribe = Softdescribe
		company.SoftCleardata = Softcleardata
		company.RomVersion = Romversion
		company.RomUpdateUrl = Romupdateurl
		company.RomMd5 = Rommd5
		company.RomDescribe = Romdescribe
		company.RomCleardata = Romcleardata
		company.ApkVersion = Apkversion
		company.ApkUpdateUrl = Apkupdateurl
		company.ApkMd5 = Apkmd5
		company.ApkDescribe = Apkdescribe
		company.ApkClerdata = Apkcleardata
		company.CustomerUrl = Customerurl
		company.ImagesUrl = Imagesurl
		company.WechatDataUrl = Wechatdataurl
		company.UserMaxCount = Usermaxcount
		company.TotalMaxCount = Totalmaxcount
		company.FuncCustomer = FunCustomer
		company.FuncImage = FuncImage
		company.FuncUploadWechat = FuncUploadWechat
		company.FuncLimitCustomer = FuncLimitCustomer
		company.FuncAutoAdd = FuncAutoAdd
		company.FuncAutoSearch = FuncAutoSearch
		company.FuncLocation = FuncLocation
		company.FuncShake = FuncShake
		company.CusttomerPush = CusttomerPush
		company.CustomerUser = CustomerUser
		_, err := m.UpdateCompany(&company)
		if err != nil {
			this.Alert("修改用户信息失败", "../company/list")
		}
		this.Alert("修改用户信息成功", "../company/list")
	} else {
		this.Data["company"] = company
		this.TplName = "company/edit.html"
	}
}

//删除公司信息
func (this *CompanyController) DeleteCompany() {
	Id, _ := this.GetInt64("Id")
	_, err := m.DelCompanyById(Id)
	if err != nil {
		beego.Error("delete is err", err)
	} else {
		this.Alert("删除成功", "/hnacenter/company/list")
	}
}

//更改公司状态
func (this *CompanyController) ChangeStatus() {
	Id, _ := this.GetInt64("Id")
	//根据id取得公司信息
	var StatusId int64
	var temnum int
	company := m.GetCompanyById(Id)
	for _, companys := range company {
		StatusId = companys["Status"].(int64)
		if StatusId == 1 {
			temnum = 0
			m.ChangeStatusById(Id, temnum)
			companys["Statusname"] = "关闭"
		} else if StatusId == 0 {
			temnum = 1
			m.ChangeStatusById(Id, temnum)
			companys["Statusname"] = "正常"
		}
	}
	this.Ctx.Redirect(302, "/hnacenter/company/list")
}

//根据组合条件进行搜索,models层中间不是采用关系查询,
func (this *CompanyController) SelectCompany() {
	provice := this.GetString("province")
	conditions := this.GetString("conditions")
	filters := this.GetString("filters")
	var status int
	if provice == "全国" {
		if filters == "公司代码" {
			company := m.GetCompanyInfoByCode(conditions)
			this.Data["json"] = &company
		} else if filters == "公司名" {
			company := m.GetCompanyByCompanyName(conditions) //模糊查询暂时还没有做
			this.Data["json"] = &company
		} else if filters == "状态" {
			if conditions == "正常" {
				status = 1
				company := m.GetCompanyByStatus(status)
				this.Data["json"] = &company
			} else if conditions == "关闭" {
				status = 0
				company := m.GetCompanyByStatus(status)
				this.Data["json"] = &company
			}
		}
	} else {
		if filters == "公司代码" {
			company := m.GetCompanyByProvinceAndCode(provice, conditions)
			this.Data["json"] = company
		} else if filters == "公司名" {
			company := m.GetCompanyByProvinceAndName(provice, conditions)
			this.Data["json"] = &company
		} else if filters == "状态" {
			if conditions == "正常" {
				status = 1
				companys := m.GetCompanyByProvinceAndStatus(provice, status)
				this.Data["json"] = &companys
			} else if conditions == "关闭" {
				status = 0
				companys := m.GetCompanyByProvinceAndStatus(provice, status)
				this.Data["json"] = &companys
			}

		}
	}
	this.TplName = "company/index.html"
}
