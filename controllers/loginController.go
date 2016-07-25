package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	m "hnacenter/models"
)

type LoginController struct {
	beego.Controller
}

// 返回一个 弹窗
func (this *LoginController) Alert(msg string, localurl string) {
	alert := fmt.Sprintf("<script>alert('" + msg + "');location.href='" + localurl + "';</script>")
	this.Ctx.WriteString(alert)
}

func (this *LoginController) Result(status bool, info string) {
	this.Data["json"] = &map[string]interface{}{"status": status, "info": info}
	this.ServeJSON()
}

func (this *LoginController) ToLogin() {
	this.TplName = "login.html"
}

func (this *LoginController) Login() {
	isAjax := this.GetString("isAjax")
	if isAjax == "0" {
		username := this.GetString("username")
		password := this.GetString("password")
		user, err := m.CheckLogin(username, password)
		if err == nil {
			this.SetSession("user", user)
			this.Ctx.Redirect(302, "/index")
			return
		} else {
			this.Data["err"] = err
			this.TplName = "err/error404.html"
			return
		}
	}
	userInfo := this.GetSession("user")
	if userInfo != nil {
		this.Ctx.Redirect(302, "login.html")
	} else {
		this.TplName = "login.html"
	}
}

// 退出
func (this *LoginController) Logout() {
	this.DelSession("user")
	this.DelSession("resource")
	this.Ctx.Redirect(302, "login.html")
}
