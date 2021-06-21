package main


import (
	"fmt"
	"unsafe"
)

/*
任何类型的 *T 都可以转换为 unsafe.Pointer；

unsafe.Pointer 也可以转换为任何类型的 *T；

unsafe.Pointer 可以转换为 uintptr；

uintptr 也可以转换为 unsafe.Pointer。
*/

type person struct {
	Name string
	Age  int
}

func main() {
	i:= 10
	ip:=&i // 取i的地址
	var fp = (*float64)(unsafe.Pointer(ip))
	*fp = *fp * 3
	fmt.Println(i)
}

func edit() {
	p := new(person)
	//Name是person的第一个字段不用偏移，即可通过指针修改

	pName := (*string)(unsafe.Pointer(p))
	*pName = "飞雪无情"

	//Age并不是person的第一个字段，所以需要进行偏移，这样才能正确定位到Age字段这块内存，才可以正确的修改
	pAge := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + unsafe.Offsetof(p.Age)))
	*pAge = 20
	fmt.Println(*p)
}
