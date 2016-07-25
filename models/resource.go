package models

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	//tool "ihytrade/src/tool"
)

type Resource struct {
	Id         int64  // 标识
	Reskey     string `orm:"size(64)" form:"Reskey" valid:"Required"`                       // 资源Key
	Level      int64  `orm:"default(1);size(11)" form:"Level" valid:"Required"`             // 资源级别
	Fid        int64  `orm:"size(11)"  form:"Fid" valid:"Required"`                         // 资源父标识
	Url        string `orm:"size(64)" form:"Url"  valid:"Required"`                         // 资源链接
	Status     int64  `orm:"size(64);default(0)" form:"Status" valid:"Required;Range(1,2)"` // 0启用 1隐藏
	Sort       int64  `orm:"size(64)" form:"Sort" valid"Required"`                          // 资源排序
	Ico        string `orm:"size(64)" form:"Ico" valid"Required"`                           // 资源图标
	Isfunction int64  `orm:"size(64)" form:"Isfunction valid"Required"`                     // 是否功能

}

func init() {
	orm.RegisterModel(new(Resource))
}
func (r *Resource) TableName() string {
	return "resource"
}

//验证数据
func (r *Resource) Checkuser() (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(&r)
	if !b {
		for _, err := range valid.Errors {
			return errors.New(err.Message)
		}
	}
	return nil
}

//根据角色id查询所有的资源
func GetResourceTree(roleId int64) (resource []orm.Params, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("role_resources").Filter("role_id", roleId).Values(&resource)
	if err != nil {
		return resource, err
	}
	return resource, err
}

//根据id查询资源
func GetResourceById(id int64) (resource Resource) {
	o := orm.NewOrm()
	resources := Resource{Id: id}
	err := o.Read(&resources, "Id")
	if err == orm.ErrNoRows {
		beego.Error("can't find the Resource", err)
	} else if err == orm.ErrMissPK {
		beego.Error("can't find the PK", err)
	}
	return resources
}

//分类
func GetTreeAndLv(resource []Resource, fid int64, lv int64) []Resource {
	var tree []Resource = make([]Resource, 0)
	for _, v := range resource {
		if v.Fid == fid {
			v.Level = lv
			child := GetTreeAndLv(resource, v.Id, lv+1)
			tree = append(tree, v)
			tree = append(tree, child...)
		}
	}
	return tree
}

//获取资源
func GetTree(isFunction int64) []Resource {
	o := orm.NewOrm()
	var resources []Resource
	_, err := o.QueryTable("resource").Filter("IsFunction", isFunction).All(&resources)
	if err != nil {
		return resources
	}
	return resources
}

//取得所有的资源
func GetAllTree() []Resource {
	o := orm.NewOrm()
	var resources []Resource
	_, err := o.QueryTable("resource").All(&resources)
	if err != nil {
		return resources
	}
	return resources
}

//根据用户信息取得用户所能访问的所有资源的url集合
func GetResourceUrlByUser(userRid int64) []string {
	o := orm.NewOrm()
	var resources []orm.Params
	var resourceId int64
	var listUrl []string
	_, err := o.QueryTable("role_resources").Filter("role_id", userRid).Values(&resources) //根据Rid取得一个用户可以访问的所有的资源id
	if err != nil {
		beego.Error("no resources can get", err)
	} else {
		for _, v1 := range resources {
			resourceId = v1["Resource"].(int64)
			resource := GetResourceById(resourceId) //循环取得所有的资源然后取得url放进数组
			if resource.Level == 1 && resource.Fid == 0 {
			} else {
				listUrl = append(listUrl, resource.Url)
			}
		}
	}
	return listUrl
}
