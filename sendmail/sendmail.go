/*
发送获取到的节点信息到目标邮箱
20180304 by panndora 死猪
*/

package sendmail

import (
	"gopkg.in/gomail.v2"
)

func Send(from string, pws string, to string, dv string) (err error) {

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "最新双S列表信息")
	m.SetBody("text/html", dv)

	d := gomail.NewDialer("smtp.qq.com", 465, from, pws)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
		return err
	}
	return nil
}
