package controllers

import (
	"github.com/astaxie/beego"
	m "hnacenter/models"
	utils "hnacenter/src/utils"
	"os"
	"time"
)

type CustomerController struct {
	LoginController
}

//资源导入第一步
func (this *CustomerController) SourceImport() {
	id := this.GetString("id")
	if id == "kv-1" {
		this.Ctx.WriteString("content")
	} else {

		this.Data["TypeList"] = m.CustomerTypeList()
		this.TplName = "source/sourceimport.html"
	}
}

// 资源上传第二步
func (this *CustomerController) SourceImportTwo() {
	_, fh, err := this.GetFile("file")
	_type, _ := this.GetInt64("type")

	if err != nil {
		beego.Error(err)
	} else {

		// nowtime := utils.Format(time.Now(), "yyyy-MM-dd")
		dest := "./static/customer/"
		os.MkdirAll(dest, 0755)
		filename := utils.Format(time.Now(), "yyyy-MM-dd-HH-mm-ss") + fh.Filename
		newpath := dest + filename
		this.SaveToFile("file", newpath)

		userinfo := this.GetSession("user").(m.User)
		RepeatNum, FailNum, SuccessNum, _ := m.CustomerReadLine(newpath, userinfo.Id, _type)

		this.Data["RepeatNum"] = RepeatNum
		this.Data["FailNum"] = FailNum
		this.Data["SuccessNum"] = SuccessNum
		this.TplName = "source/sourceimporttwo.html"
	}

}

//分配资源
func (this *CustomerController) DistributionSource() {
	provices := m.GetAllProvice()
	this.Data["provice"] = provices
	this.TplName = "source/disbutionsource.html"
}

//公司列表
func (this *CustomerController) CompanysList() {
	if this.IsAjax() {
		sEcho := this.GetString("sEcho")
		iStart, _ := this.GetInt64("iDisplayStart")
		iLength, _ := this.GetInt64("iDisplayLength")
		companys, count := m.GetConpanyList(iStart, iLength, "Id")
		data := make(map[string]interface{})
		data["aaData"] = companys
		data["iTotalDisplayRecords"] = count
		data["iTotalRecords"] = iLength
		data["sEcho"] = sEcho
		this.Data["json"] = &data
		this.ServeJSON()
	} else {
		this.Data["ActionUrl"] = beego.AppConfig.String("ActionUrl")
		this.TplName = "source/compnaylist.html"
	}
}

//资源查询
func (this *CustomerController) SourceSelect() {
	if this.IsAjax() {
		sEcho := this.GetString("sEcho")
		iStart, _ := this.GetInt64("iDisplayStart")
		iLength, _ := this.GetInt64("iDisplayLength")
		companys, count := m.GetConpanyList(iStart, iLength, "Id")

		data := make(map[string]interface{})
		data["aaData"] = companys
		data["iTotalDisplayRecords"] = count
		data["iTotalRecords"] = iLength
		data["sEcho"] = sEcho
		this.Data["json"] = &data
		this.ServeJSON()
	} else {
		this.Data["ActionUrl"] = beego.AppConfig.String("ActionUrl")
		this.TplName = "source/sourceselect.html"
	}
}

//分发日志
func (this *CustomerController) DistributionLog() {
	this.TplName = "source/distributionrecord.html"
}

//分发记录
func (this *CustomerController) DistributionRecord() {
	this.TplName = "source/distributionrecord.html"
}
