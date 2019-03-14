package basic

import (
	"fmt"
	"os"
	"os/exec"
)

// init 包中的init会被首先调用，用于初始化
func init() {
	fmt.Println("basic init...")
}

// OutAbsPath 输出绝对路径
func OutAbsPath() string {
	path, _ := exec.LookPath(os.Args[0])
	return path
}

// SliceTest test
func SliceTest() {
	data := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	s := data[:2:8] // 第三个参数代表cap
	fmt.Println(s)
	fmt.Println(cap(s))
}
