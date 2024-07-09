package email

import (
	"bytes"
	"codearena/pkg/utils"
	"html/template"
	"log"
	"net/smtp"
	"path/filepath"

	"github.com/jordan-wright/email"
)

var (
	EmailPool  *email.Pool
	username   = "1241236275@qq.com"
	authcode   = "dgbnheclbztdhjhi"
	from_email = "love-story <1241236275@qq.com>"
	tpl_name   = "confirm_signup"
)

func init() {
	var err error
	EmailPool, err = email.NewPool(
		"smtp.qq.com:587",
		4,
		smtp.PlainAuth("", username, authcode, "smtp.qq.com"),
	)

	if err != nil {
		log.Fatalf("failed to create pool:%v\n", err.Error())
	}
}

func SendEmailWithTmpl(to string, code string) error {
	e := email.NewEmail()
	e.Sender = "LoveStory"
	e.From = from_email
	e.To = []string{to}
	e.Subject = "感谢注册 LoveStory 应用！下面是您的专属验证码："

	// 读取confirm_signup.gohtml文件
	rootDir, _ := utils.GetRootDir()
	filePath := filepath.Join(rootDir, "backend", "resources", tpl_name+".gohtml")
	tpl, err := template.ParseFiles(filePath)
	if err != nil {
		log.Printf("failed to parse template file:%v\n", err.Error())
		return err
	}

	// 定义一个结构体来包装code变量
	type EmailData struct {
		Code string
	}
	data := EmailData{Code: code}

	// 生成html内容，发送邮件
	var body bytes.Buffer
	if err := tpl.Execute(&body, data); err != nil {
		log.Printf("failed to execute template:%v\n", err.Error())
		return err
	}
	e.HTML = body.Bytes()
	if err := EmailPool.Send(e, -1); err != nil {
		log.Printf("failed to send email:%v\n", err.Error())
		return err
	}

	return nil
}
