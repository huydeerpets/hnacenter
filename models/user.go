package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"hnacenter/src/tool"
)

type User struct {
	Id       int64
	Usercode string `orm:"not null;unique" form:"Usercode" valid:"Required"` //用户代码
	Username string `orm:"not null;unique" form:"Uname" valid:"Required"`    //用户名称
	Pwd      string `orm:"not null;size(64)" form:"Pwd" valid:"Required"`    //密码
	Rid      int64  `orm:"not null; form:"Rid" valid:"Required"`             //角色id
	// Oid      int64  `orm:"not null;size(64)" form:"Oid"  valid:"Required"`                //机构ID
	Status int64   `orm:"size(11);default(0)" form:"Status" valid:"Required;Range(1,2)"` //状态 0正常 1，禁用
	Remark string  `orm:"null;size(200)" form:"Remark" valid:"MaxSize(200)"`             //备注
	Role   []*Role `orm:"rel(m2m)"`
}

func init() {
	orm.RegisterModel(new(User))
}

func (u *User) TableName() string {
	return "user"
}

//表单数据验证
func checkUser(u *User) (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(&u)
	if !b {
		for _, err := range valid.Errors {
			return errors.New(err.Message)
		}
	}
	return nil
}

// 登录检查
func CheckLogin(username string, pwd string) (user User, err error) {
	user = GetUserByUsername(username)
	if user.Id == 0 {
		return user, errors.New("user is not exist")
	}
	password := tool.EncodeUserPwd(pwd)
	if user.Pwd != password {
		return user, errors.New("password is wrong ")
	}
	return user, nil
}

//根据用户名查找用户
func GetUserByUsername(username string) (user User) {
	o := orm.NewOrm()
	users := User{Username: username}
	err := o.Read(&users, "Username")
	if err == orm.ErrNoRows {
		beego.Error("can't find the User", err)
	} else if err == orm.ErrMissPK {
		beego.Error("can't find the PK", err)
	}
	return users
}

//根据角色id得到资源id
func GetResourceByRid(rid int64) (resources []orm.Params, err error) {
	o := orm.NewOrm()
	num, err := o.QueryTable("role_resources").Filter("role_id", rid).Values(&resources)
	fmt.Println(num)
	return resources, err

}

//查询所有的用户
func GetAllUser(istart int64, ilength int64, id string) (users []orm.Params, count int64) {
	o := orm.NewOrm()
	user := new(User)
	qs := o.QueryTable(user)
	qs.Limit(ilength, istart).OrderBy(id).Values(&users)
	count, _ = qs.Count()
	return users, count
}

//根据id查询用户
func GetUserById(id int64) (user User) {
	o := orm.NewOrm()
	users := User{Id: id}
	err := o.Read(&users, "Id")
	if err == orm.ErrNoRows {
		beego.Error("can't find the user", err)
	} else if err == orm.ErrMissPK {
		beego.Error("can't find the PK", err)
	}
	return users
}

//修改用户
func UpdateUser(u *User) (int64, error) {
	if err := checkUser(u); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	user := make(orm.Params)
	if len(u.Username) > 0 {
		user["Username"] = u.Username
	}
	if len(u.Usercode) > 0 {
		user["Usercode"] = u.Usercode
	}
	if u.Status != 0 {
		user["Status"] = u.Status
	}
	user["Pwd"] = u.Pwd
	user["Rid"] = u.Rid
	user["Remark"] = u.Remark
	if len(user) == 0 {
		return 0, errors.New("Update field ")
	}
	var table User
	num, err := o.QueryTable(table).Filter("Id", u.Id).Update(user)
	return num, err
}

//增加用户
func AddNewUser(u *User) (int64, error) {
	if err := checkUser(u); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	users := new(User)
	users.Usercode = u.Usercode
	users.Username = u.Username
	users.Pwd = u.Pwd
	users.Remark = u.Remark
	users.Rid = u.Rid
	users.Status = u.Status
	id, err := o.Insert(users)
	if err != nil {
		beego.Error("insert data into user is err", err)
	}
	return id, err
}

//根据Id号删除用户
func DelUser(id int64) (int64, error) {
	o := orm.NewOrm()
	user := User{Id: id}
	status, err := o.Delete(&user)
	return status, err
}
