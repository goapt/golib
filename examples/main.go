package main

import (
	"net/url"
	"fmt"
)

func main()  {

	c := "https://pay.verystar.cn/admin/submerchant/index?status=&category=&parent_user_id=&provider_id=&bank_user_id=&mch_alipay=&user_name=600220"
	u,_ := url.Parse(c)
	fmt.Println(u.String())
}
