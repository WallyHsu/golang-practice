package main

import (
	"fmt"
	"strconv"

	"./wallypackage"
)

/*
	程式入口
*/
func main() {
	var a int = 1
	var b int32 = 2
	var c int64 = 3
	var d string = "999"
	var e float32 = 88.8
	var f float64 = 99.9
	var x string = "I Love 859"

	//題目1
	g := int32(a) + b
	fmt.Printf("題目1: a + b = %d\n", g)

	//題目2
	h := int64(a) + int64(b) + c
	fmt.Printf("題目2: a + b + c = %d\n", h)

	//題目3
	i := f / float64(e)
	fmt.Printf("題目3: f / e = %f\n", i)

	//題目4
	o := a + strToInt(d)
	fmt.Printf("題目4: a + d = %d\n", o)

	//題目5
	x += wallypackage.IntToStr(a)
	fmt.Printf("題目5: 字串相接 = %s\n", x)
}

//題目6
func strToInt(str string) int {
	res, _ := strconv.Atoi(str)
	return res
}
