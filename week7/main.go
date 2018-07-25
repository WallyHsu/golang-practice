package main

import (
	"./package"
)

type Htc struct {
	Name string
}

func (h Htc) Sensor() string {
	return h.Name
}
func (h Htc) Monitor() string {
	return h.Name
}
func (h Htc) Button() string {
	return h.Name
}

type Apple struct {
	Name string
}

func (a Apple) Sensor() string {
	return a.Name
}
func (a Apple) Monitor() string {
	return a.Name
}
func (a Apple) Button() string {
	return a.Name
}

func main() {
	htc := Htc{Name: "M10"}
	apple := Apple{Name: "Iphone X"}

	phone.Lock(htc)
	phone.Unlock(htc)
	phone.Call(htc)

	println()

	phone.Lock(apple)
	phone.Unlock(apple)
	phone.Call(apple)
}
