/*
该程序用于每日自动登录瓜瓜ss签到并获取最新的节点列表信息发送到邮箱中待用
20180304 by deadpig panndora 死猪
*/
package main

import (
	"log"
	"os"
	"strings"

	"github.com/Oguangpan/owscs/laodconfig"
	"github.com/Oguangpan/owscs/sendmail"
	"github.com/Oguangpan/owscs/signin"
)

// 配置文件存放地址
const confFile string = "./config.cfg"

func main() {
	cfg := new(laodconfig.Configs)
	logFile, err := os.Create("owscs.log")
	defer logFile.Close()
	if err != nil {
		log.Fatalln("无法打开log文件!")
	}
	runLog := log.New(logFile, "[OWSCS运行日志]", log.LstdFlags)

	if err := cfg.Laod(confFile); err != nil {
		runLog.Println("导入配置文件错误,", err)
		return
	}
	runLog.Println("读取配置文件成功")

	node, err := signin.LoginAndSignin(cfg.LoginName, cfg.LoginPws, cfg.Hosturl)
	if err != nil {
		runLog.Println("错误", err)
		return
	}
	runLog.Println("签到成功,正在发送邮件")

	//b := strings.Join(node, "<p>")
	b := strings.Join(node, ",<br>")

	if err := sendmail.Send(cfg.SendMail, cfg.SendPws, cfg.Mail, b); err != nil {
		runLog.Println("邮件发送失败,原因是: ", err)
		return
	}
	runLog.Println("邮件发送成功,程序正常退出")
	return
}
