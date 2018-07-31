package main

import (
	"fmt"
	"time"
)

type SmartPhone interface {
	Name() string
	Size() int
	EnduranceTime() time.Duration
}

type Iphone struct {
	version       string
	width, height int
	battery       time.Duration
}

func (i Iphone) Name() string {
	return i.version
}

func (i Iphone) Size() int {
	return i.height * i.width
}

func (i Iphone) EnduranceTime() time.Duration {
	return i.battery * time.Hour
}

type Note struct {
	version       string
	width, height int
	battery       time.Duration
}

func (p Note) Name() string {
	return p.version
}

func (p Note) Size() int {
	return p.height * p.width
}

func (p Note) EnduranceTime() time.Duration {
	return p.battery * time.Minute
}

//他是圓形的
type IWatch struct {
	version string
	strap   int
	battery time.Duration
}

func (i IWatch) Name() string {
	return i.version
}

func (i IWatch) Size() int {
	return i.strap
}

func (i IWatch) EnduranceTime() time.Duration {
	return i.battery * time.Hour
}

func main() {
	iphone := Iphone{width: 30, height: 90, battery: 24, version: "iphone-8"}
	note := Note{width: 40, height: 120, battery: 1300, version: "note-3"}
	iw := IWatch{strap: 10, battery: 500, version: "Series-2"}

	ShowPhone(iphone)
	ShowPhone(pixel)
	ShowPhone(iw)
}

func ShowPhone(s SmartPhone) {
	fmt.Printf("Product Name: %v \n", s.Name())
	fmt.Printf("Size: %v \n", s.Size())
	fmt.Printf("Endurance Time: %v \n", s.EnduranceTime())
	fmt.Println()
}
