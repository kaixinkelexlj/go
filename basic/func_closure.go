package basic

import "fmt"

func ClosureTest() {
	c := a()
	c()
	c()
	c()
}

func a() func() int {
	i := 0
	b := func() int {
		i++
		fmt.Println(i)
		return i
	}
	return b
}
