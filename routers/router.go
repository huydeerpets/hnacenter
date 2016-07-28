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
	beego.Router("/hnacenter/company/list", &controllers.CompanyController{}, "*:GetCompanyList")       //公司列表
	beego.Router("/hnacenter/company/add", &controllers.CompanyController{}, "*:AddCompany")            //增加公司
	beego.Router("/hnacenter/company/edit", &controllers.CompanyController{}, "*:EditCompany")          //修改公司
	beego.Router("/hnacenter/company/delete", &controllers.CompanyController{}, "*:DeleteCompany")      //删除公司
	beego.Router("/hnacenter/company/changestatus", &controllers.CompanyController{}, "*:ChangeStatus") //更改状态
	beego.Router("/hnacenter/company/select", &controllers.CompanyController{}, "*:SelectCompany")
	beego.Router("/hnacenter/companyconfig/add", &controllers.CompanyConfigController{}, "*:AddCompanyConfig")
	beego.Router("/hnacenter/company/companychange", &controllers.CompanyController{}, "*:ChangeCompanyAttr") //资源管理中间的公司列表的属性修改

	//设备管理路由
	beego.Router("/hnacenter/device/list", &controllers.DeviceInfoController{}, "*:GetDeviceLists") //设备列表
	beego.Router("/hnacenter/device/delete", &controllers.DeviceInfoController{}, "*:DelDevice")    //删除设备

	//资源管理路由
	beego.Router("/hnacenter/source/importsource", &controllers.CustomerController{}, "*:SourceImport")       //资源导入
	beego.Router("/hnacenter/source/importtwo", &controllers.CustomerController{}, "*:SourceImportTwo")       //资源导入2
	beego.Router("/hnacenter/source/distribution", &controllers.CustomerController{}, "*:DistributionSource") //分配资源
	beego.Router("/hnacenter/source/companylist", &controllers.CustomerController{}, "*:CompanysList")        //公司列表
	beego.Router("/hnacenter/source/select", &controllers.CustomerController{}, "*:SourceSelect")             //资源查询
	beego.Router("/hnacenter/source/logging", &controllers.CustomerController{}, "*:DistributionLog")         //分发日志
	beego.Router("/hnacenter/source/record", &controllers.CustomerController{}, "*:DistributionRecord")       //分发记录

	//资源分组
	beego.Router("/hnacenter/source/grouping", &controllers.UserCustomerGroupController{}, "*:SourceDistribution") //资源分组
	beego.Router("/hnacenter/source/addusergroup", &controllers.UserCustomerGroupController{}, "*:AddUserGroup")   //添加用户资源组
}
