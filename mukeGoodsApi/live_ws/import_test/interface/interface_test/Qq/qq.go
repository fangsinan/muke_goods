package Qq

import "fmt"

type QqAdmin struct {
	Name string
}

func (n *QqAdmin) GetUserId() int {
	fmt.Println("this is Qq  get admin id ")
	fmt.Println(n.Name)
	return 1
}
