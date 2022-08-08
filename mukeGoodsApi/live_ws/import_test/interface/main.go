package main

import (
	"fmt"
)

type Phone interface {
	call()
}

type NokiaPhone struct {
	Phone Phone
	Name  string
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
	fmt.Println(nokiaPhone.Name)
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func main() {

	no := &NokiaPhone{
		Name: "as",
	}
	no.call()

	ip := &IPhone{}
	ip.call()

}
