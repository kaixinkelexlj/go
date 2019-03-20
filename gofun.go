package main

import "fmt"
import "./fun"
import basic "./basic"

func main() {
	fmt.Println("go fun")
	fun.Ok()
	fmt.Println(basic.OutAbsPath())
	basic.SliceTest()
	basic.TestMap()
	basic.PointerCalculate()
	basic.PointerCalSettingValue()
}
