package controllers

import (
	// "github.com/astaxie/b/eego"
	// "github.com/astaxie/beego/orm"
	m "hnacenter/models"
	// "reflect"/
)

type IndexController struct {
	ResourceController
}

// 跳转到主页面
func (this *IndexController) ToIndex() {
	user := this.GetSession("user")
	this.Data["user"] = user
	tree := this.GetTree()
	this.Data["tree"] = tree
	users := user.(m.User)
	var resourceId int64
	resources, _ := m.GetResourceByRid(users.Rid)
	if resources != nil {
		res := make([]m.Resource, len(resources))
		for k1, v1 := range resources {
			resourceId = v1["Resource"].(int64)
			resource := m.GetResourceById(resourceId)
			res[k1].Id = resource.Id
			res[k1].Fid = resource.Fid
			res[k1].Reskey = resource.Reskey
			res[k1].Ico = resource.Ico
			res[k1].Level = resource.Level
			res[k1].Url = resource.Url
			res[k1].Status = resource.Status
			res[k1].Sort = resource.Sort
			res[k1].Isfunction = resource.Isfunction
		}
		this.SetSession("resource", res)
	}

	this.TplName = "index.html"
	this.Layout = "common/layout.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["HtmlHead"] = "common/html_head.html"
	this.LayoutSections["HtmlLeft"] = "common/html_left.html"
	this.LayoutSections["HtmlFooter"] = "common/html_footer.html"

}
func (this *IndexController) ToMain() {
	this.TplName = "index.html"
}
