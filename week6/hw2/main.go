package main

import (
	"fmt"

	"./package"
)

type Rocket struct {
	Name string
	Fuel int
}

func (r Rocket) Flight() {
	fmt.Println(r.Name + " 航行中!!!")
}

func main() {
	rocket1 := Rocket{Name: "肥鼠號", Fuel: 3}
	rocket2 := Rocket{Name: "哈士奇號", Fuel: 3}
	rocket.Start(rocket1)
	rocket.Start(rocket2)
}
