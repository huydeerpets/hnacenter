package controllers

import (
	"github.com/astaxie/beego"
	m "hnacenter/models"
)

type CompanyConfigController struct {
	LoginController
}

//增加启动配置信息
func (this *CompanyConfigController) AddCompanyConfig() {
	isAction := this.GetString("isAction")
	if isAction == "0" {
		Code := this.GetString("Code")
		DatabseInfo := this.GetString("DatabseInfo")
		Addr := this.GetString("Addr")
		Httpport, _ := this.GetInt("Httpport")
		DayMaxDownloadCount, _ := this.GetInt("DayMaxDownloadCount")
		DayMaxHidDownload, _ := this.GetInt("DayMaxHidDownload")
		ClearDataTimes := this.GetString("ClearDataTimes")
		Status, _ := this.GetInt("Status")
		ServerAddr := this.GetString("ServerAddr")
		Ormdebug, _ := this.GetInt("Ormdebug")
		Testdebug, _ := this.GetInt("Testdebug")
		Maxuploadimages, _ := this.GetInt("Maxuploadimages")
		Adminuser, _ := this.GetInt64("Adminuser")
		Recycle, _ := this.GetInt64("Recycle")
		//根据Code取出配置信息
		companyconfig := m.GetConfigInfo(Code)
		if companyconfig.Id == 0 {
			companyconfig.Code = Code
			companyconfig.DatabseInfo = DatabseInfo
			companyconfig.Addr = Addr
			companyconfig.Httpport = Httpport
			companyconfig.DayMaxDownloadCount = DayMaxDownloadCount
			companyconfig.DayMaxHidDownload = DayMaxHidDownload
			companyconfig.ClearDataTimes = ClearDataTimes
			companyconfig.Status = Status
			companyconfig.ServerAddr = ServerAddr
			companyconfig.Ormdebug = Ormdebug
			companyconfig.Testdebug = Testdebug
			companyconfig.Maxuploadimages = Maxuploadimages
			companyconfig.Adminuser = Adminuser
			companyconfig.Recycle = Recycle
			_, err := m.AddCompanyconfigInfo(&companyconfig)
			if err != nil {
				beego.Error("insert is err", err)
				return
			}
		} else {
			this.Alert("添加成功", "")
		}
	}
}
