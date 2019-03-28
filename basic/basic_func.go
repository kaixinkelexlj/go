package basic

import (
	"fmt"
	"os"
	"os/exec"
	"unsafe"

	"github.com/axgle/mahonia"
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

// TestMap test
func TestMap() {

	m := map[string]string{"key0": "value0", "key1": "value1"}
	fmt.Printf("map m : %v\n", m)
	//map插入
	m["key2"] = "value2"
	fmt.Printf("inserted map m : %v\n", m)
	//map修改
	m["key0"] = "hello world!"
	fmt.Printf("updated map m : %v\n", m)
	//map查找
	val, ok := m["key0"]
	if ok {
		fmt.Printf("map's key0 is %v\n", val)
	}

	// 长度：获取键值对数量。
	len := len(m)
	fmt.Printf("map's len is %v\n", len)

	// cap 无效，error
	// cap := cap(m)    //invalid argument m (type map[string]string) for cap
	// fmt.Printf("map's cap is %v\n", cap)

	// 判断 key 是否存在。
	if val, ok = m["key"]; !ok {
		fmt.Println("map's key is not existence")
	}

	// 删除，如果 key 不存在，不会出错。
	if val, ok = m["key1"]; ok {
		delete(m, "key1")
		fmt.Printf("deleted key1 map m : %v\n", m)
	}

}

// 将 Pointer 转换成 uintptr，可变相实现指针运算。
func PointerCalculate() {
	d := struct {
		s string
		x int
	}{"abc", 100}

	p := uintptr(unsafe.Pointer(&d)) // *struct -> Pointer -> uintptr
	p += unsafe.Offsetof(d.x)        // uintptr + offset

	p2 := unsafe.Pointer(p) // uintptr -> Pointer
	px := (*int)(p2)        // Pointer -> *int
	*px = 200               // d.x = 200

	fmt.Printf("%#v\n", d)
}

func PointerCalSettingValue() {

	type user struct {
		name string
		age  int
	}

	u := new(user)
	fmt.Println(*u)

	pName := (*string)(unsafe.Pointer(u))
	*pName = "张三"

	pAge := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(u)) + unsafe.Offsetof(u.age)))
	*pAge = 20

	fmt.Println(*u)

}

func ConvertToString(src string, srcCode string, dstCode string) string {
	scrCoder := mahonia.NewDecoder(srcCode)
	scrString := scrCoder.ConvertString(src)
	_, dstString, _ := mahonia.NewDecoder(dstCode).Translate([]byte(scrString), true)
	return string(dstString)
}

func TestPanic() {
	func() {
		defer func() {
			if err := recover(); err != nil {
				println(err.(string)) // 将 interface{} 转型为具体类型。
			}
		}()
		panic("panic error!")
	}()
}

func Try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fun()
}
