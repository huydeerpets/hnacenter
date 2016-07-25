package autoschudle

// 定时凌晨查询资源剩余，改变公司状态
import (
	"github.com/astaxie/beego"
	"github.com/bitly/go-simplejson"
	m "hnnacenter/models"
	"hnnacenter/src/jobs"
	utils "hnnacenter/src/utils"

	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/smtp"
	"strconv"
	"strings"
)

const (
	user     = "1121163988@qq.com"
	password = "lc13135926589"
	host     = "smtp.qq.com:25"
	to       = "coco@touzi.org"
)

type job struct {
}

func (j *job) Run() {
	var (
		warningerr     []string
		warningcompany []string
	)

	companylist, _, _ := m.GetFieldCompanyNum(1)
	warning := new(m.CustomerWarning)
	for _, item := range companylist {
		push := item["CusttomerPush"].(int64)
		if push == 1 {
			num, err := GetCompanyRemain(item["Code"].(string))
			if err != nil {
				warning.Code = item["Code"].(string)
				warning.Remain = num
				warning.Warning = err.Error()
				warning.Types = 1
				inserid, err := m.InsertCustomerWarning(warning)
				if err != nil {
					beego.Error(inserid, err)
				}
				warningerr = append(warningerr, item["Name"].(string))
			} else {
				if num < 5000 {
					warning.Code = item["Code"].(string)
					warning.Remain = num
					warning.Types = 2
					inserid, err := m.InsertCustomerWarning(warning)
					if err != nil {
						beego.Error(inserid, err)
					}
					warningcompany = append(warningcompany, item["Name"].(string))
				}
			}
		}
	}

	if len(warningcompany) > 0 || len(warningerr) > 0 {
		subject := "资源警告邮寄通知"
		body := fmt.Sprintf(`<html><body><h3>"%s 资源告急"<br>"%s 公司异常"</h3></body></html>`, strings.Join(warningcompany, ","), strings.Join(warningerr, ","))
		err := SendMail(user, password, host, to, subject, body, "html")
		if err != nil {
			beego.Error(err.Error())
		}
	}
}

func AutoSchudePush() {
	j := &job{}
	taskName := "autoschudepush"
	jobs.Schedule("0 28 17 * * *", j, taskName)
}

// 请求公司剩余资源数量
func GetCompanyRemain(company string) (int64, error) {
	companyconfig := m.GetAllCompanyConfig(company)
	serveraddr := companyconfig.ServerAddr + ":" + strconv.Itoa(companyconfig.Httpport) + "/api"
	body := fmt.Sprintf(`{"action":"remaincustomer","checkstr":"powerbyihaoyue","company":"%s"}`, company)
	dataStrEncopy := utils.MainEncrypt(body)
	parameter := map[string]string{"data": dataStrEncopy}
	res, reserr := httpnRmageRequest(serveraddr, parameter)
	if reserr != nil {
		return -1, reserr
	} else {
		result, err := ioutil.ReadAll(res.Body)
		if err != nil { //读取返回Body错误
			return -1, err
		} else {
			query, err := utils.MainDecrypt(string(result))
			if err != nil { //解密失败
				return -1, err
			} else {
				js, err := simplejson.NewJson([]byte(query))
				if err != nil { //JSON解析失败
					return -1, err
				} else {
					remain := js.Get("remain").MustInt64()
					return remain, nil
				}
			}
		}
	}
}

// Creates a new file upload http request with optional extra params
func httpnRmageRequest(uri string, params map[string]string) (*http.Response, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}

	content_type := writer.FormDataContentType()

	err := writer.Close()
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("POST", uri, body)

	req.Header.Set("Content-Type", content_type)
	resp, err := http.DefaultClient.Do(req)
	return resp, err
}

// 发送邮寄
func SendMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}
