package models

import (
	"errors"
	// "fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

type Role struct {
	Id       int64
	Name     string      `orm:"size(64)" form:"Name" valid:"Required" `            //角色名
	Rolekey  string      `orm:"size(64)" form:"Rolekey" valid:"Required"`          //角色key
	Remark   string      `orm:"null;size(200)" form:"Remark" valid:"MaxSize(200)"` //备注
	Resource []*Resource `orm:"rel(m2m)"`
	// User     []*User     `orm:"reverse(many)"`
}

func init() {
	orm.RegisterModel(new(Role))
}

func (r *Role) TableName() string {
	return "role"
}

//检查role表
func checkRole(r *Role) (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(&r)
	if !b {
		for _, err := range valid.Errors {
			return errors.New(err.Message)
		}
	}
	return nil
}

//根据角色id查找角色
func GetRoleByRid(rid int64) (role Role) {
	o := orm.NewOrm()
	roles := Role{Id: rid}
	err := o.Read(&roles, "Id")
	if err == orm.ErrNoRows {
		beego.Error("can't find this role", err)
	} else if err == orm.ErrMissPK {
		beego.Error("can't find the PK", err)
	}
	return roles
}

//角色列表
func GetRoleList() []Role {
	o := orm.NewOrm()
	var roles []Role
	_, err := o.QueryTable("role").All(&roles)
	if err != nil {
		return roles
	}
	return roles
}

//查找出所有的角色
func GetRList(start int64, length int64, sort string) (roles []orm.Params, count int64) {
	o := orm.NewOrm()
	role := new(Role)
	qs := o.QueryTable(role)
	qs.Limit(length, start).OrderBy(sort).Values(&roles)
	count, _ = qs.Count()
	return roles, count
}

//根据角色名取得role
func GetRole(name string) (role Role) {
	o := orm.NewOrm()
	roles := Role{Name: name}
	err := o.Read(&roles, "Name")
	if err == orm.ErrNoRows {
		beego.Error("can't find the role", err)
	} else if err == orm.ErrMissPK {
		beego.Error("can't find the PK", err)
	}
	return roles
}

//根据Id删除角色
func DelRole(id int64) (int64, error) {
	o := orm.NewOrm()
	role := Role{Id: id}
	status, err := o.Delete(&role)
	return status, err
}

//根据名字取得角色
func GetRoleByName(name string) (role Role) {
	o := orm.NewOrm()
	roles := Role{Name: name}
	err := o.Read(&roles, "Name")
	if err == orm.ErrNoRows {
		beego.Error("can't find the role", err)
	} else if err == orm.ErrMissPK {
		beego.Error("can't find the PK", err)
	}
	return roles
}

//增加角色
func Addrole(r *Role) (int64, error) {
	if err := checkRole(r); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	role := new(Role)
	role.Name = r.Name
	role.Remark = r.Remark
	role.Rolekey = r.Rolekey
	id, err := o.Insert(role)
	if err != nil {
		beego.Error("insert into data is err", err)
	}
	return id, err
}

//根据id取得角色
func GetRoleById(id int64) (role Role) {
	o := orm.NewOrm()
	roles := Role{Id: id}
	err := o.Read(&roles, "Id")
	if err == orm.ErrNoRows {
		beego.Error("can't find the role", err)
	} else if err == orm.ErrMissPK {
		beego.Error("can't find the PK", err)
	}
	return roles
}

//更新角色
func Updaterole(r *Role) (int64, error) {
	if err := checkRole(r); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	role := make(orm.Params)
	if len(r.Name) > 0 {
		role["Name"] = r.Name
	}
	if len(r.Remark) > 0 {
		role["Remark"] = r.Remark
	}
	if len(r.Rolekey) > 0 {
		role["Rolekey"] = r.Rolekey
	}
	if len(role) == 0 {
		return 0, errors.New("Update field")
	}
	var table Role
	num, err := o.QueryTable(table).Filter("Id", r.Id).Update(role)
	return num, err
}

//根据角色id删除角色资源
func DelRoleResourceByRoleId(roldId int64) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("role_resources").Filter("role_id", roldId).Delete()
	return err
}

//给角色增加资源
func AddRoleResource(roleId int64, rId int64) (int64, error) {
	o := orm.NewOrm()
	role := Role{Id: roleId}
	resource := Resource{Id: rId}
	m2m := o.QueryM2M(&role, "Resource")
	num, err := m2m.Add(&resource)
	return num, err
}
