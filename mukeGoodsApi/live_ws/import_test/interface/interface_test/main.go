package main

import (
	"fmt"
	"webApi/live_ws/import_test/interface/interface_test/Weixin"
)

type Admin interface {
	GetUserId() int
}

type UserServe struct {
	Admin Admin
}

func main() {
	u := UserServe{

		//Admin: &Qq2.QqAdmin{
		//	Name: "222",
		//},
		Admin: &Weixin.WeixinAdmin{
			Name: "微信admin",
		},
	}
	id := u.Admin.GetUserId()
	fmt.Println(id)

}
