package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	// "github.com/astaxie/beego/orm"
	m "hnacenter/models"
	_ "hnacenter/routers"
	"strings"
)

//验证用户登录
var FilterUser = func(ctx *context.Context) {
	user := ctx.Input.Session("user")
	if user == nil {
		ctx.Redirect(302, "/login")
	} else {
		_, ok := user.(m.User)
		if !ok && ctx.Request.RequestURI != "/login" {
			ctx.Redirect(302, "/login")
		}
	}

}

//界面权限认证(非管理员用户，取得可以访问的所有url，然后取得界面url，然后进行比较，有则可以访问，没有就不可以访问)
var Check = func(ctx *context.Context) {
	admingUser := beego.AppConfig.String("admin_user")
	userinfo := ctx.Input.Session("user")
	userRid := userinfo.(m.User).Rid
	ActionUrl := beego.AppConfig.String("ActionUrl")
	if userinfo == nil {
		ctx.Redirect(302, ActionUrl)
	} else {
		if userinfo.(m.User).Username == admingUser {
			return
		} else {
			URList := m.GetResourceUrlByUser(userRid) //取得用户能访问的资源url集合
			//用户请求的uri,数据类型是string
			requestURL := strings.Split(ctx.Request.RequestURI, "?")
			requestNewURL := requestURL[0]
			for i := 0; i < len(URList); i++ {
				if requestNewURL == URList[i] {
					return
				}
			}
			ctx.Redirect(302, ActionUrl)
		}

	}
}

func main() {
	beego.InsertFilter("/hnacenter/*", beego.BeforeRouter, FilterUser)
	beego.InsertFilter("/hnacenter/*", beego.BeforeRouter, Check)
	beego.Run()

}
