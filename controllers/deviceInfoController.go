package controllers

import (
	"github.com/astaxie/beego"
	m "hnacenter/models"
)

type DeviceInfoController struct {
	LoginController
}

//设备列表
func (this *DeviceInfoController) GetDeviceLists() {
	if this.IsAjax() {
		sEcho := this.GetString("sEcho")
		iStart, _ := this.GetInt64("iDisplayStart")
		iLength, _ := this.GetInt64("iDisplayLength")
		device, count := m.GetDeviceList(iStart, iLength, "Id")
		for _, devices := range device {
			switch devices["Status"] {
			case int64(0):
				devices["Statusname"] = "正常"
			case int64(1):
				devices["Statusname"] = "设备未审核"
			default:
				devices["Statusname"] = "异常"
			}
		}
		data := make(map[string]interface{})
		data["aaData"] = device
		data["iTotalDisplayRecords"] = count
		data["iTotalRecords"] = iLength
		data["sEcho"] = sEcho
		this.Data["json"] = &data
		this.ServeJSON()
	} else {
		this.Data["ActionUrl"] = beego.AppConfig.String("ActionUrl")
		this.TplName = "deviceinfo/index.html"
	}
}

//删除设备信息
func (this *DeviceInfoController) DelDevice() {
	if this.IsAjax() {
		Id, _ := this.GetInt64("Id")
		_, err := m.DelDeviceById(Id)
		if err == nil {
			this.Result(true, "删除成功")
		} else {
			this.Result(false, "删除失败")
		}
	} else {
		this.TplName = "/"
	}
}
