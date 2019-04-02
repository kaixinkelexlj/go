package main

import (
	"fmt"

	basic "./basic"
	jsonx "./jsonfun"
)

func main() {
	/* fmt.Println("go fun")
	fun.Ok()
	fmt.Println(basic.OutAbsPath())
	basic.SliceTest()
	basic.TestMap()
	basic.PointerCalculate()
	basic.PointerCalSettingValue()
	basic.ClosureTest()
	basic.TestPanic()
	basic.Try(func() {
		panic("test panic")
	}, func(err interface{}) {
		fmt.Println(err)
	})*/
	u := basic.User{1, "xulujun"}
	u.GetUserNameAndPrint()
	err := basic.OpenFile("not_exists")
	switch v := err.(type) {
	case *basic.PathError:
		fmt.Println(" get path error, ", v)
	default:
	}
	basic.OpenFile("/Users/didi/login.sh")
	basic.Reflect()
	basic.DoTestReflectStruct()
	jsonx.TestStructJSON()
}
