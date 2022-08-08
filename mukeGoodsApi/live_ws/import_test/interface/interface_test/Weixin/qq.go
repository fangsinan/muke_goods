package Weixin

import "fmt"

type WeixinAdmin struct {
	Name string
}

func (n *WeixinAdmin) GetUserId() int {
	fmt.Println("this is 微信  get 微信 admin id ")
	fmt.Println(n.Name)
	return 1
}
