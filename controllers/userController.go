package controllers

import (
	"github.com/astaxie/beego"
	m "hnacenter/models"
)

type UserController struct {
	LoginController
}

//用户列表
func (this *UserController) GetUserList() {
	user := this.GetSession("user")
	if this.IsAjax() {
		sEcho := this.GetString("sEcho")
		iStart, _ := this.GetInt64("iDisplayStart")
		iLength, _ := this.GetInt64("iDisplayLength")

		users, count := m.GetAllUser(iStart, iLength, "Id")
		for _, user := range users {
			switch user["Status"] {
			case int64(0):
				user["Statusname"] = "启用"
			case int64(1):
				user["Statusname"] = "禁用"
			default:
				user["Statusname"] = "账号异常"

			}
			role := m.GetRoleByRid(user["Rid"].(int64))
			user["RoleName"] = role.Name

		}
		data := make(map[string]interface{})
		this.Data["user"] = user
		data["aaData"] = users
		data["iTotalDisplayRecords"] = count
		data["iTotalRecords"] = iLength
		data["sEcho"] = sEcho
		this.Data["json"] = &data
		this.ServeJSON()
	} else {
		this.Data["ActionUrl"] = beego.AppConfig.String("ActionUrl")
		this.Data["user"] = user
		this.TplName = "user/index.html"
	}
}

//更新用户
func (this *UserController) UpdateUsers() {
	isAction := this.GetString("isAction")
	Id, _ := this.GetInt64("Id")
	Username := this.GetString("Username")
	Usercode := this.GetString("Usercode")
	Pwd := this.GetString("Pwd")
	Status, _ := this.GetInt64("Status")
	Remark := this.GetString("Remark")
	RoleName := this.GetString("selectRole")
	roless := m.GetRole(RoleName)
	Rid := roless.Id
	user := m.GetUserById(Id)
	// 获取角色列表
	roles := m.GetRoleList()
	this.Data["json"] = roles
	var roleName []string
	for i := 0; i < len(roles); i++ {
		roname := roles[i].Name
		roleName = append(roleName, roname)
	}
	this.Data["RoleNameList"] = roleName
	role := m.GetRoleByRid(user.Rid)
	this.Data["RoleName"] = role.Name
	if "0" == isAction {
		user.Username = Username
		user.Usercode = Usercode
		user.Pwd = Pwd
		user.Rid = Rid
		user.Status = Status
		user.Remark = Remark
		_, err := m.UpdateUser(&user)
		if err != nil {
			this.Alert("修改用户信息失败", "../user/list")
		}
		this.Alert("修改用户信息成功", "../user/list")
	} else {
		this.Data["user"] = user
		this.TplName = "user/edit.html"
	}
}

//增加用户
func (this *UserController) AddUser() {
	isAction := this.GetString("isAction")
	//取得所有的角色名
	roles := m.GetRoleList()
	this.Data["roles"] = roles
	if isAction == "0" {
		Username := this.GetString("Username")
		Usercode := this.GetString("Usercode")
		Pwd := this.GetString("Pwd")
		roleNames := this.GetString("addRole")
		roless := m.GetRole(roleNames)
		Rid := roless.Id
		Status, _ := this.GetInt64("Status")
		Remark := this.GetString("Remark")

		user := m.GetUserByUsername(Username)
		if user.Id == 0 {
			user.Pwd = Pwd
			user.Remark = Remark
			user.Rid = Rid
			user.Status = Status
			user.Usercode = Usercode
			user.Username = Username
			_, err := m.AddNewUser(&user)
			if err != nil {
				beego.Error("insert new user is err", err)
				return
			}
		} else {
			this.Alert("添加成功", "../user/list")
		}
	} else {

		this.TplName = "user/add.html"
	}

}

//删除用户
func (this *UserController) DeleteUser() {
	Id, _ := this.GetInt64("Id")
	_, err := m.DelUser(Id)
	if err != nil {
		this.Alert("删除用户信息失败", "../user/list")
	} else {
		this.Alert("删除用户信息成功", "../user/list")
	}
}
