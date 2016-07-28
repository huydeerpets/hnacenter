package controllers

import (
	"github.com/astaxie/beego"
	m "hnacenter/models"
	"strconv"
	"strings"
)

type RoleController struct {
	LoginController
}

//获取角色列表
func (this *RoleController) GetRoleList() {
	user := this.GetSession("user")
	if this.IsAjax() {
		sEcho := this.GetString("sEcho")
		iStart, _ := this.GetInt64("iDisplayStart")
		iLength, _ := this.GetInt64("iDisplayLength")
		roles, count := m.GetRList(iStart, iLength, "Id")
		data := make(map[string]interface{})
		data["aaData"] = roles
		data["iTotalDisplayRecords"] = count
		data["iTotalRecords"] = iLength
		data["sEcho"] = sEcho
		this.Data["json"] = &data
		this.ServeJSON()
	} else {
		this.Data["ActionUrl"] = beego.AppConfig.String("ActionUrl")
		this.Data["user"] = user
		this.TplName = "role/index.html"
	}

}

//更新角色信息
func (this *RoleController) UpdateRoles() {
	isAction := this.GetString("isAction")
	Id, _ := this.GetInt64("Id")
	role := m.GetRoleById(Id)
	this.Data["role"] = role
	if isAction == "0" {
		Name := this.GetString("Name")
		Rolekey := this.GetString("Rolekey")
		Remark := this.GetString("Remark")
		role.Name = Name
		role.Rolekey = Rolekey
		role.Remark = Remark
		_, err := m.Updaterole(&role)
		if err != nil {
			beego.Error("Update is err", err)
			return
		} else {
			this.Alert("修改成功", "../role/list")
		}
	}
	this.TplName = "role/edit.html"

}

//增加角色
func (this *RoleController) AddRole() {
	isAction := this.GetString("isAction")
	if isAction == "0" {
		Name := this.GetString("Name")
		Rolekey := this.GetString("Rolekey")
		Remark := this.GetString("Remark")
		role := m.GetRoleByName(Name)
		if role.Id == 0 {
			role.Name = Name
			role.Rolekey = Rolekey
			role.Remark = Remark
			_, err := m.Addrole(&role)
			if err != nil {
				beego.Error("insert new role is err", err)
				return
			}
		} else {
			this.Alert("添加成功", "../role/list")
		}
	} else {

		this.TplName = "role/add.html"
	}

}

//删除角色
func (this *RoleController) DeleteRole() {
	Id, _ := this.GetInt64("Id")
	_, err := m.DelRole(Id)
	if err != nil {
		this.Alert("删除用户信息失败", "../role/list")
	} else {
		this.Alert("删除用户信息成功", "../role/list")
	}
}

//拿到资源列表
func (this *RoleController) DisPermission() {
	//取得角色id
	Id, _ := this.GetInt64("Id")
	//通过id取得角色
	role := m.GetRoleByRid(Id)
	//取得所有的资源
	resource := m.GetAllTree()
	//给资源分级
	tree := m.GetTreeAndLv(resource, 0, 1)
	this.Data["role"] = role
	this.Data["json"] = tree
	this.TplName = "role/permission.html"
}

//设置权限
func (this *RoleController) AddPermission() {
	Ids := this.GetString("Ids")
	Id, _ := this.GetInt64("Id")
	//根据角色Id删除角色不需要的资源
	errId := m.DelRoleResourceByRoleId(Id)
	if errId != nil {
		beego.Error("delete role_resources is failure", errId)
	}
	for _, v := range strings.Split(Ids, ",") {
		rId, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			beego.Error("type is err", err)
		}
		_, adderr := m.AddRoleResource(Id, rId)
		if adderr != nil {
			this.Result(false, "失败")
		} else {
			this.Result(true, "成功")
		}
	}
}
