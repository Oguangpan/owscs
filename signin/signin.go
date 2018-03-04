package signin

import (
	"regexp"

	"github.com/yudeguang/gather"
)

func LoginAndSignin(n string, p string, u string) (r []string, err error) {
	postdata := map[string]string{"email": n, "passwd": p, "remember_me": "week"}
	gat := gather.NewGather("chrome", false)

	var nodeHtml string
	u1, u2, u3 := u+"/user/_login.php", u+"/user/_checkin.php", u+"/user/node.php"
	// 打开页面
	_, rurl, err := gat.Get(u, u)
	if err != nil {
		return nil, err
	}
	// 模拟登录
	_, rurl, err = gat.Post(u1, rurl, postdata)
	if err != nil {
		return nil, err
	}
	// 点击签到
	_, rurl, err = gat.Get(u2, rurl)
	if err != nil {
		return nil, err
	}
	// 获取列表
	nodeHtml, rurl, _ = gat.Get(u3, rurl)

	reg, e := regexp.Compile("node_json\\.php\\?id\\={1}\\d+")
	if e != nil {
		return nil, e
	}
	var nodelist []string
	for _, v := range reg.FindAllString(nodeHtml, -1) {

		html, _, _ := gat.Get(u+"/user/"+v, rurl)
		nodelist = append(nodelist, html)
	}

	return nodelist, nil
}
