package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	tool "hnacenter/src/utils"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

//设备信息表
type DeviceInfo struct {
	Id           int64
	Davice       string `orm:"index"` //设备号
	RegisterTime int64  //注册时间
	Status       int    `orm:"default(0)"` //设备状态(0:正常 1：设备未审核)
	UserId       string //注册设备用户
	Serial       string //Serial(CA699HWKIMIKRFRK
	Deviceid     string //设备ID(865369029028431)
	MacAddr      string //Mac地址(18:dc:56:73:38:8a)
	BassAddr     string //无线地址(20:dc:e6:0e:dd:7e)
	Ssid         string //服务集标识(ChinaNet-HR8U)
	Android      string //安卓id(b2ad8ea67a3a687c)
	Imsi         string //国际移动用户识别码(460022093117918)
	Iccid        string //集成电路卡识别
	PhoneNumber  string //手机号码
	Company      string //公司代码
	IpAddress    string //IP地址
	DeviceType   int    `orm:"default(0);"` //设备类型  0为模拟器 1位手机
	// Phonemodel   int    `orm:"default(0);"` // 手机型号 0红米1s
}

func (this *DeviceInfo) TableName() string {
	return "device_info"
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(DeviceInfo))
}

func AddDevice(username string, hid string, company string, ipaddress string) (id int64, device DeviceInfo, err error) {
	PhoneNumber := CheckDeviceMobile()
	model := orm.NewOrm()

	device.Davice = hid
	device.RegisterTime = time.Now().Unix()
	device.Status = 0
	device.UserId = username
	device.Serial = CheckDeviceSerial()
	device.Deviceid = CheckDeviceImei()
	device.MacAddr = CheckDeviceMacAddr()
	device.BassAddr = CheckDeviceBassAddr()
	device.Ssid = CheckDeviceSsid()
	device.Android = CheckDeviceAndriod()
	device.Imsi = CheckDeviceImsi(PhoneNumber)
	device.Iccid = CheckDeviceIccid(PhoneNumber)
	device.PhoneNumber = PhoneNumber
	device.Company = company
	device.IpAddress = ipaddress

	id, err = model.Insert(&device)
	return id, device, err
}

//获取设备信息
func ReadDeviceInfo(hid string, username string, company string) (*DeviceInfo, error) {
	model := orm.NewOrm()
	fmt.Println(hid)
	var deviceinfo DeviceInfo
	err := model.QueryTable("device_info").Filter("davice", hid).Filter("user_id", username).One(&deviceinfo)
	return &deviceinfo, err
}

// 用户名的设备帐号数量
func DeviceUnameNum(uname string, company string) (DeviceInfo []orm.Params, DeviceCount int64) {
	modal := orm.NewOrm()
	DeviceCount, err := modal.QueryTable("device_info").Filter("UserId", uname).Filter("Company", company).Values(&DeviceInfo)
	if err != nil {
		beego.Error("没有账户存在")
	}
	return DeviceInfo, DeviceCount
}

// 公司名下的设备数量
func DeviceAllCompanyNum(company string) (DeviceInfo []orm.Params, DeviceCount int64) {
	model := orm.NewOrm()
	DeviceCount, err := model.QueryTable("device_info").Filter("Company", company).Values(&DeviceInfo)
	if err != nil {
		beego.Error("没有账户存在")
	}
	return DeviceInfo, DeviceCount
}

// 检查device是否正确
func CheckDevice(device string) bool {
	if len(device) == 36 {
		reg, err := regexp.Compile("[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}")
		if reg.MatchString(device) {
			return true
		} else {
			beego.Error(err)
			return false
		}
	}
	return false
}

// 检查SERIAL[CA699HWKIXFOCZWL]
func CheckDeviceSerial() string {
	var info string = ""
	str := "CA699HWKI" + tool.RandomStringSpec1(7, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	if ReadDevice("Serial", str) {
		info = CheckDeviceSerial()
	} else {
		info = str
	}
	return info

}

// 检查IMEI[865369029941047]
func CheckDeviceImei() string {
	var info string = ""
	imeisub := "865369029" + tool.RandomNumeric(5)
	imei := imeisub + GetLastImei(imeisub)
	if ReadDevice("Deviceid", imei) {
		info = CheckDeviceImei()
	} else {
		info = imei
	}
	return info
}

// 校验IMEI最后一位
// IMEI校验码算法：
// (1).将偶数位数字分别乘以2，分别计算个位数和十位数之和
// (2).将奇数位数字相加，再加上上一步算得的值
// (3).如果得出的数个位是0则校验位为0，否则为10减去个位数
func GetLastImei(imei string) string {
	// var str string
	var (
		sumou    int    = 0
		sumji    int    = 0
		trueimei string = ""
	)
	for i := 0; i < 14; i++ {
		if i%2 != 0 {
			d, _ := strconv.Atoi(string(imei[i]))
			evennum := strconv.Itoa(d * 2)
			if len(evennum) == 2 {
				e1, _ := strconv.Atoi(string(evennum[0]))
				e2, _ := strconv.Atoi(string(evennum[1]))
				sumou = sumou + e1 + e2
			} else {
				e, _ := strconv.Atoi(evennum)
				sumou = sumou + e
			}
		} else {
			d, _ := strconv.Atoi(string(imei[i]))
			sumji = sumji + d
		}
	}
	sum := strconv.Itoa(sumou + sumji)
	number := string(sum[len(sum)-1])
	if number == "0" {
		trueimei = "0"
	} else {
		lastimei, _ := strconv.Atoi(number)
		id := 10 - lastimei
		trueimei = strconv.Itoa(id)
	}
	return trueimei
}

// MacAddr地址生成[18:dc:56:cc:b6:18]
func CheckDeviceMacAddr() string {
	var info string = ""
	s := "0123456789abcdef"
	str := "18:dc:56:" + tool.RandomStringSpec1(2, s) + ":" + tool.RandomStringSpec1(2, s) + ":" + tool.RandomStringSpec1(2, s)
	if ReadDevice("MacAddr", str) {
		info = CheckDeviceMacAddr()
	} else {
		info = str
	}
	return info
}

// BassAddr地址生成[20:dc:e6:7a:f2:f7]
func CheckDeviceBassAddr() string {
	var info string = ""
	s := "0123456789abcdef"
	str := "20:dc:e6:" + tool.RandomStringSpec1(2, s) + ":" + tool.RandomStringSpec1(2, s) + ":" + tool.RandomStringSpec1(2, s)
	if ReadDevice("BassAddr", str) {
		info = CheckDeviceBassAddr()
	} else {
		info = str
	}
	return info
}

// Ssid生成[ChinaNet-MRHY]
func CheckDeviceSsid() string {
	var info string = ""
	s := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	str := "ChinaNet-" + tool.RandomStringSpec1(4, s)
	if ReadDevice("Ssid", str) {
		info = CheckDeviceSsid()
	} else {
		info = str
	}
	return info
}

// Android生成[b2ad8ea661688d9e]
func CheckDeviceAndriod() string {
	var info string = ""
	s := "0123456789abcdef"
	str := "b2ad8ea6" + tool.RandomStringSpec1(8, s)
	if ReadDevice("Android", str) {
		info = CheckDeviceAndriod()
	} else {
		info = str
	}
	return info
}

// Imsi生成[]4600220 9087 1143
// 截取 4600220 phone[3:4] 1143;
func CheckDeviceImsi(phone string) string {
	var info string = ""
	subphone := phone[3:7]
	str := "4600220" + subphone + tool.RandomNumeric(4)
	if ReadDevice("Imsi", str) {
		info = CheckDeviceImsi(phone)
	} else {
		info = str
	}
	return info
}

// Iccid生成 898600 9 1 7 14 032243670
//898600 letter subphone sub3 tool.RandomNumeric(1) 14 tool.RandomNumeric(8)
func CheckDeviceIccid(phone string) string {
	var info string = ""
	letter := "0123456789abcdef"
	subphone := phone[3:4]
	sub3 := rand.Intn(4)
	str := "898600" + tool.RandomStringSpec1(1, letter) + subphone + strconv.Itoa(sub3) + tool.RandomNumeric(1) + "14" + tool.RandomNumeric(8)
	if ReadDevice("Iccid", str) {
		info = CheckDeviceIccid(phone)
	} else {
		info = str
	}
	return info
}

//检查手机号码是否唯一
func CheckDeviceMobile() string {
	var info string = ""
	phone := []string{"1709273", "1709311", "1709087", "1700901"}
	p := string(phone[rand.Intn(4)]) + tool.RandomNumeric(4)
	if ReadDevice("PhoneNumber", p) {
		info = CheckDeviceMobile()
	} else {
		info = p
	}
	return info
}

// 数据库查询
func ReadDevice(filed string, value string) bool {
	model := orm.NewOrm()
	count, _ := model.QueryTable("device_info").Filter(filed, value).Count()
	fmt.Println(count)
	if count > 0 {
		return true
	} else {
		return false
	}
}

// 查询设备的信息
func GetDeviceInfoByHid(hid string) (d DeviceInfo) {
	model := orm.NewOrm()
	model.QueryTable("device_info").Filter("Davice", hid).One(&d)
	return d

}

//获取设备列表
func GetDeviceList(page int64, page_size int64, sort string) (devices []orm.Params, count int64) {
	omodel := orm.NewOrm()

	device := new(DeviceInfo)
	qs := omodel.QueryTable(device)
	var cond *orm.Condition
	cond = orm.NewCondition()

	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.SetCond(cond).Limit(page_size, offset).OrderBy(sort).Values(&devices)
	count, _ = qs.SetCond(cond).Count()
	return devices, count
}

//删除设备信息
func DelDeviceById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&DeviceInfo{Id: Id})
	return status, err
}

//删除设备信息
func DelDeviceByHid(company string, hid string) (int64, error) {
	o := orm.NewOrm()
	status, err := o.QueryTable("device_info").Filter("Davice", hid).Filter("Company", company).Delete()
	return status, err
}

// 转移帐号下的设备

func TransferUserDevice(company string, device string, user string, transferuser string) (int64, error) {
	model := orm.NewOrm()
	num, err := model.QueryTable("device_info").Filter("Company", company).Filter("Davice", device).Filter("UserId", user).Update(orm.Params{
		"UserId": transferuser,
	})
	if err != nil {
		return 0, err
	}
	return num, err
}

//添加手机设备
func AddMobileDevice(device DeviceInfo) (id int64, d DeviceInfo, err error) {
	PhoneNumber := CheckDeviceMobile()
	model := orm.NewOrm()

	device.RegisterTime = time.Now().Unix()
	device.Status = 0
	if device.Serial == "" {
		device.Serial = CheckDeviceSerial()
	}
	if device.Deviceid == "" {
		device.Deviceid = CheckDeviceImei()
	}
	if device.MacAddr == "" {
		device.MacAddr = CheckDeviceMacAddr()
	}
	if device.BassAddr == "" {
		device.BassAddr = CheckDeviceBassAddr()
	}
	if device.Ssid == "" {
		device.Ssid = CheckDeviceSsid()
	}
	if device.Android == "" {
		device.Android = CheckDeviceAndriod()
	}
	if device.PhoneNumber == "" {
		device.PhoneNumber = PhoneNumber
	}
	if device.Imsi == "" {
		device.Imsi = CheckDeviceImsi(device.PhoneNumber)
	}
	if device.Iccid == "" {
		device.Iccid = CheckDeviceIccid(device.PhoneNumber)
	}

	id, err = model.Insert(&device)
	return id, device, err
}

// 查询设备总数
func GetTotalDeviceNum() (int64, error) {
	model := orm.NewOrm()
	count, err := model.QueryTable("device_info").Count()
	return count, err
}

//根据设备状态查询设备数量
func GetNumByStatus(status int) (int64, error) {
	o := orm.NewOrm()
	count, err := o.QueryTable("device_info").Filter("Status", status).Count()
	return count, err
}
