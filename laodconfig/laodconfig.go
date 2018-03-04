/*
读取同目录下的配置文件.
20180304 by panndora 死猪
*/
package laodconfig

import (
	"github.com/robfig/config"
)

type Configs struct {
	LoginName string
	LoginPws  string
	Hosturl   string
	LoginUrl  string
	SigninUrl string
	NodeUrl   string
	SendMail  string
	SendPws   string
	Mail      string
}

func (p *Configs) Laod(confFile string) (err error) {
	cfg, err := config.ReadDefault(confFile)
	if err != nil {
		return err
	}
	p.Hosturl, _ = cfg.String("info", "hosturl")
	p.LoginName, _ = cfg.String("info", "name")
	p.LoginPws, _ = cfg.String("info", "pws")
	p.SendMail, _ = cfg.String("info", "sendmail")
	p.SendPws, _ = cfg.String("info", "sendpws")
	p.Mail, _ = cfg.String("info", "tomail")
	return nil
}
