package routers

import (
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/orm"
	"hnacenter/controllers"
	m "hnacenter/models"
	. "hnacenter/src/tool"
	"mime"
)

func init() {
	initialize()
	router()
	beego.AddFuncMap("stringsToJson", StringtoJson)
}

func initialize() {
	mime.AddExtensionType(".css", "text/css")

	//链接数据库
	m.InitDB()
}

func router() {
	//登录界面路由
	beego.Router("/", &controllers.LoginController{}, "*:ToLogin")
	//取得账户名和密码然后进入主界面的路由
	beego.Router("/login", &controllers.LoginController{}, "*:Login")
	//登出
	beego.Router("/logout", &controllers.LoginController{}, "*:Logout")
	beego.Router("/index", &controllers.IndexController{}, "*:ToIndex")
	beego.Router("/main", &controllers.IndexController{}, "*:ToMain")
	//用户
	beego.Router("/hnacenter/user/list", &controllers.UserController{}, "*:GetUserList")
	beego.Router("/hnacenter/user/edit", &controllers.UserController{}, "*:UpdateUsers")
	beego.Router("/hnacenter/user/add", &controllers.UserController{}, "*:AddUser")
	beego.Router("/hnacenter/user/del", &controllers.UserController{}, "*:DeleteUser")
	//角色
	beego.Router("/hnacenter/role/list", &controllers.RoleController{}, "*:GetRoleList")
	beego.Router("/hnacenter/role/edit", &controllers.RoleController{}, "*:UpdateRoles")
	beego.Router("/hnacenter/role/add", &controllers.RoleController{}, "*:AddRole")
	beego.Router("/hnacenter/role/del", &controllers.RoleController{}, "*:DeleteRole")
	//角色权限路由
	beego.Router("/hnacenter/role/permission", &controllers.RoleController{}, "*:AddPermission")
	beego.Router("/hnacenter/role/setPermission", &controllers.RoleController{}, "*:DisPermission")
	//公司管理路由
	beego.Router("/hnacenter/company/list", &controllers.CompanyController{}, "*:GetCompanyList")
	beego.Router("/hnacenter/company/add", &controllers.CompanyController{}, "*:AddCompany")
	beego.Router("/hnacenter/company/edit", &controllers.CompanyController{}, "*:EditCompany")
	beego.Router("/hnacenter/company/delete", &controllers.CompanyController{}, "*:DeleteCompany")
	beego.Router("/hnacenter/company/changestatus", &controllers.CompanyController{}, "*:ChangeStatus")
	beego.Router("/hnacenter/company/select", &controllers.CompanyController{}, "*:SelectCompany")

	//设备管理路由
	beego.Router("/hnacenter/device/list", &controllers.DeviceInfoController{}, "*:GetDeviceLists")
	beego.Router("/hnacenter/device/delete", &controllers.DeviceInfoController{}, "*:DelDevice")
}
