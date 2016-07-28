package controllers

import (
	"github.com/astaxie/beego"
	m "hnacenter/models"
)

type UserCustomerGroupController struct {
	LoginController
}

//添加资源用户组
func (this *UserCustomerGroupController) AddUserGroup() {
	isAction := this.GetString("isAction")
	if isAction == "0" {
		userCustomerGroup := this.GetString("uCustomerGroup")
		person := this.GetString("selected_list")
		usergroup := m.GetUserGroupByName(userCustomerGroup)
		if usergroup.Id == 0 {
			usergroup.Name = userCustomerGroup
			usergroup.UserGroup = person
			_, err := m.AdduCustomerGroup(&usergroup)
			if err != nil {
				beego.Error("UserCustomerGroupController insert is err", err)
			}
		} else {
			this.Alert("添加成功", "../source/grouping")
		}
		this.Ctx.Redirect(302, "/hnacenter/source/addusergroup")
	} else {
		this.TplName = "source/sourcegrouptwo.html"
	}
}

//资源分组列表
func (this *UserCustomerGroupController) SourceDistribution() {
	if this.IsAjax() {
		sEcho := this.GetString("sEcho")
		iStart, _ := this.GetInt64("iDisplayStart")
		iLength, _ := this.GetInt64("iDisplayLength")
		usergroup, count := m.GetuCustomerGroupList(iStart, iLength, "Id")
		data := make(map[string]interface{})
		data["aaData"] = usergroup
		data["iTotalDisplayRecords"] = count
		data["iTotalRecords"] = iLength
		data["sEcho"] = sEcho
		this.Data["json"] = &data
		this.ServeJSON()
	} else {
		this.Data["ActionUrl"] = beego.AppConfig.String("ActionUrl")
		this.TplName = "source/sourcegroup.html"

	}

}
