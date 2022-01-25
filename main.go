package main

import (
	"fmt"
	"golem/controller"
	"golem/model"
)

func main() {
	fmt.Println("hello world")
	name := controller.Hellonew("quyet")
	fmt.Println(name)
	fmt.Println(model.Cong(1, 2))
	fmt.Println("hello")
	fmt.Println("quyet")
	fmt.Println("vip")
}
