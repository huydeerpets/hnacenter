package models

import (
	"bufio"
	// "errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

//客户资源表
type Customer struct {
	Id         int64
	Name       string        //姓名
	Mobile     int64         `orm:"unique"`     //手机号
	UserId     int64         `orm:"index"`      //用户ID
	Distribute int           `orm:"default(0)"` //是否下发 0：没有下发 1：下发
	BatchId    int64         `orm:"index"`      //资源导入的批次
	AreaId     int64         `orm:"index"`      //资源所属的区域区号
	AreaName   string        `orm:"index"`      //资源所属的区域名称
	Operators  string        // 运营商
	Province   int64         `orm:"index"` //省份代码
	Sendtime   int64         //分发时间
	Company    string        `orm:"index"`              //分发公司
	Type       *CustomerType `orm:"rel(fk);default(1)"` //资源类型
	Sendfail   int           `orm:"index;default(0)"`   //发送失败 0 未分发/分发失败 1 分发成功
}

func (this *Customer) TableName() string {
	return "customer"
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Customer))
}

// 导入资源
func CustomerReadLine(filePth string, username int64, _type int64) (int64, int64, int64, error) {
	model := orm.NewOrm()
	var RepeatNum int64 = 0  //重复数量
	var FailNum int64 = 0    //不合格数量
	var SuccessNum int64 = 0 //成功数量
	var EmptyNum int64 = 0   //空行数量
	cdid, err := AddCustomerDis(username, _type)
	if err != nil {
		beego.Error(err)
		return RepeatNum, FailNum, SuccessNum, err
	}
	f, err := os.Open(filePth)
	if err != nil {
		beego.Error(err)
		return RepeatNum, FailNum, SuccessNum, err
	}
	defer func() {
		f.Close()
		err = os.Remove(filePth)
		if err != nil {
			beego.Error(err)
		}
	}()
	// defer f.Close()
	bfRd := bufio.NewReader(f)

	err = model.Begin()
	for {
		line, _, err := bfRd.ReadLine()
		if err == io.EOF {
			break
		}
		linedata := strings.TrimSpace(string(line))
		phone, _ := strconv.ParseInt(linedata, 10, 64)
		if len(linedata) == 0 {
			EmptyNum++
		} else if phone == 0 || len(linedata) != 11 || linedata[0:1] == "0" {
			FailNum++
		} else if boo := CheckRepeatNum(phone); boo == true {
			RepeatNum++
		} else {
			mobilearea := CheckMobileArea(linedata)
			var customer Customer
			customer.Mobile = phone
			customer.UserId = username
			customer.BatchId = cdid
			customer.AreaId = mobilearea.Mobilenet
			customer.AreaName = mobilearea.Mobilearea
			customer.Operators = mobilearea.Mobileareanumber
			customer.Province = mobilearea.Province
			customer.Type = &CustomerType{Id: _type}

			_, err := model.Insert(&customer)
			if err == nil {
				SuccessNum++
			} else {
				beego.Error(err)
			}
		}

		if err != nil {
			beego.Error(err)
			break
		}
	}
	err = model.Commit()
	if err == nil {
		var d CustomerDistribute
		d.Id = cdid
		d.DistributeNum = SuccessNum + RepeatNum + FailNum
		d.SuccessNum = SuccessNum
		d.FailNum = FailNum
		d.RepeatNum = RepeatNum
		d.TotalNum, _ = CustomerTotalNum(0)
		num := UpdateDistribute(d, "DistributeNum", "SuccessNum", "FailNum", "RepeatNum", "TotalNum")
		beego.Debug(num)
	}
	return RepeatNum, FailNum, SuccessNum, nil
}

// 统计数量or 资源下发
// uname 用户名
// SearchArea 省
// SeatchAreaCity 区域
// customerallotnum 限制资源数量
// distribute 是否被下载
// ispower 是否总权限
func GetAllCustomer(uname int64, SearchArea int64, SeatchAreaCity int64, customerallotnum int64, distribute int, cusType int64) (customers []orm.Params, count int64) {
	o := orm.NewOrm()
	customer := new(Customer)
	qs := o.QueryTable(customer)

	var cond *orm.Condition
	cond = orm.NewCondition()

	if SearchArea == 1000 {
		cond = cond.And("AreaId", 0)
	} else if SeatchAreaCity != 10000 {
		cond = cond.And("AreaId", SeatchAreaCity)
	} else if SearchArea > 0 {
		cond = cond.And("Province", SearchArea)
	}
	cond = cond.And("UserId", uname)
	cond = cond.And("Distribute", distribute)
	if cusType != -1 {
		cond = cond.And("type_id", cusType)
	}

	count, err := qs.SetCond(cond).Limit(customerallotnum).OrderBy("Id").Values(&customers)
	if err != nil {
		beego.Error(err)
	}
	return customers, count
}

// 获取以及打过标记的资源数量
func GetTagCustomerNum(company string) (Num int64, err error) {
	o := orm.NewOrm()
	Num, err = o.QueryTable("customer").Filter("Company", company).Filter("Distribute", 0).Count()
	return Num, err
}

//获取打过标记的资源
func GetTagCustomer(customernum int64, company string) (customers []orm.Params, count int64, err error) {
	o := orm.NewOrm()
	count, err = o.QueryTable("customer").Filter("Company", company).Filter("Distribute", 0).Limit(customernum).Values(&customers)
	return customers, count, err
}

// 批量更新资源
// num, err := o.QueryTable("user").Filter("name", "slene").Update(orm.Params{
//     "name": "astaxie",
// })
func UpdateAllCustomer(uname int64, SearchArea int64, SeatchAreaCity int64, customerallotnum int64, cusType int64, proviceDeny []int64, company string) (int64, error) {
	o := orm.NewOrm()
	sql := "UPDATE customer SET company='" + company + "',sendtime=" + fmt.Sprintf(`%d`, time.Now().Unix()) + " WHERE user_id =" + fmt.Sprintf("%d", uname) + " AND distribute=0 AND company=''"
	if SearchArea > 0 { //省份选择之后不执行排除省份功能
		if SearchArea == 1000 {
			sql = sql + " AND area_id = 0"
		} else if SeatchAreaCity != 10000 {
			sql = sql + " AND area_id = " + fmt.Sprintf(`%d`, SeatchAreaCity)
		} else if SearchArea > 0 {
			sql = sql + " AND province = " + fmt.Sprintf(`%d`, SearchArea)
		}
	} else {
		provicenum := len(proviceDeny)
		if provicenum > 0 {
			for i := 0; i < provicenum; i++ {
				sql = sql + " AND province != " + fmt.Sprintf(`%d`, proviceDeny[i])
			}
		}
	}

	if cusType != -1 {
		sql = sql + " AND type_id = " + fmt.Sprintf(`%d`, cusType)
	}
	sql = sql + " LIMIT " + fmt.Sprintf(`%d`, customerallotnum)

	result, err := o.Raw(sql).Exec()
	if err == nil {
		num, err := result.RowsAffected()
		return num, err
	} else {
		return 0, err
	}
}

// 检查是否重复数据
func CheckRepeatNum(phone int64) bool {
	model := orm.NewOrm()
	flag := model.QueryTable("customer").Filter("Mobile", phone).Exist()
	return flag
}

// 查询资源总数量 user 0为查看所有
func CustomerTotalNum(user int64) (count int64, err error) {
	model := orm.NewOrm()
	if user > 0 {
		count, err = model.QueryTable("customer").Filter("UserId", user).Count()
	} else {
		count, err = model.QueryTable("customer").Count()
	}

	if err != nil {
		beego.Error(err)
		return 0, err
	}
	return count, nil
}

// 查询资源的剩余数量 Ispower是否查看全部 1为查看所有
func CustomerRemainNum(Uname int64, Ispower int64, distrbute int64) (int64, error) {
	var (
		count int64
		err   error
	)
	model := orm.NewOrm()

	if Ispower == 1 {
		count, err = model.QueryTable("customer").Filter("Distribute", distrbute).Count()
	} else {
		count, err = model.QueryTable("customer").Filter("UserId", Uname).Filter("Distribute", distrbute).Count()
	}

	if err != nil {
		beego.Error(err)
		return 0, err
	}
	return count, nil
}

// 查询数量
func GetAllCustomerNumber(uname int64, SearchArea int64, SeatchAreaCity int64, distribute int64, cusType int64) int64 {
	o := orm.NewOrm()
	customer := new(Customer)
	qs := o.QueryTable(customer)

	var cond *orm.Condition
	cond = orm.NewCondition()

	if SearchArea == 1000 {
		cond = cond.And("AreaId", 0)
	} else if SeatchAreaCity != 10000 {
		cond = cond.And("AreaId", SeatchAreaCity)
	} else if SearchArea > 0 {
		cond = cond.And("Province", SearchArea)
	}
	cond = cond.And("UserId", uname)
	cond = cond.And("Distribute", distribute)
	if cusType > 0 {
		cond = cond.And("type_id", cusType)
	}

	count, err := qs.SetCond(cond).Limit(-1).OrderBy("Id").Count()
	if err != nil {
		beego.Error(err)
	}
	return count
}

//将指定批次的资源更改分类
func UpdateCusTypeByBatch(id int64, _type int64) {
	o := orm.NewOrm()
	o.QueryTable("customer").Filter("BatchId", id).Update(orm.Params{"type_id": _type})
}

// 批量更新分发资源
func CustomerUpdata(dataid []int, company string) int64 {
	model := orm.NewOrm()
	customer := make(orm.Params)
	customer["Company"] = company
	customer["Distribute"] = 1
	customer["Sendtime"] = time.Now().Unix()
	customer["Sendfail"] = 1

	num, err := model.QueryTable("customer").Filter("Id__in", dataid).Update(customer)
	if err != nil {
		beego.Error(err)
		return num
	}
	return num
}

// 批量更新分发资源不成功
func CustomerUpdataFail(dataid []int) int64 {
	model := orm.NewOrm()
	customer := make(orm.Params)
	customer["Sendfail"] = 0

	num, err := model.QueryTable("customer").Filter("Id__in", dataid).Update(customer)
	if err != nil {
		beego.Error(err)
		return num
	}
	return num
}

// 获取分类资源的总数
func GetUcustomerGroupCount(uname int64, type_id int64, customertype []int, zonedata int64) (int64, error) {
	model := orm.NewOrm()
	cond := orm.NewCondition()
	if uname > 0 {
		cond = cond.And("Userid", uname)
	}
	cond = cond.And("distribute__in", customertype)
	if type_id != -1 {
		cond = cond.And("type_id", type_id)
	}
	if zonedata > 0 && type_id != -1 {
		cond = cond.And("Province", zonedata)
	}
	GroupCount, err := model.QueryTable("customer").SetCond(cond).Count()
	return GroupCount, err
}

//资源下载 uname int64, SearchArea int64, SeatchAreaCity int64, customerallotnum int64, distribute int, cusType int64
// uname 用户名
// SearchArea 省
// SeatchAreaCity 区域
// customerallotnum 限制资源数量
// distribute 是否被下载
// ispower 是否总权限
// customer["Company"] = company
// 	customer["Distribute"] = 1
// 	customer["Sendtime"] = time.Now().Unix()
// 	customer["Sendfail"] = 1
// func CustomerDownload(uname string, SearchArea string, SeatchAreaCity string, customerallotnum int64, distribute int, cusType string, company string) (list []orm.Params, count int, err error) {
// 	count = 0
// 	o := orm.NewOrm()
// 	o.Begin()
// 	sql := "SELECT id,mobile FROM customer WHERE user_id =" + fmt.Sprintf("%d", uname) + " AND distribute=0"
// 	if SearchArea == "1000" {
// 		sql = sql + " AND area_id = 0"
// 	} else if SeatchAreaCity != "10000" {
// 		sql = sql + " AND area_id = " + SeatchAreaCity
// 	} else if SearchArea > 0 {
// 		sql = sql + " AND province = " + SearchArea
// 	}
// 	if cusType != -1 {
// 		sql = sql + " AND type_id = " + cusType
// 	}

// 	_, err = o.Raw(sql).Values(&list)
// 	if err != nil {
// 		o.Rollback()
// 		return list, 0, err
// 	}
// 	customertime := time.Now().Unix()
// 	for _, item := range list {
// 		_, err = o.QueryTable("customer").Filter("Id", item["id"].(string)).Update(orm.Params{
// 			"Company":    company,
// 			"Sendtime":   customertime,
// 			"Distribute": 1,
// 			"Sendfail":   1,
// 		})
// 		if err != nil {
// 			o.Rollback()
// 			return list, 0, err
// 		}
// 		count++
// 	}
// 	o.Commit()
// 	return list, count, nil
// }
