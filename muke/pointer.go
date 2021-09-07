package main

import (
	"fmt"
)

type person struct {
	name string
	age int
}

type animal struct {
	string
	float32
}

func pointerInt()  {
	var a *int // nil is zero value
	//*a = 10 // 需要使用new 分配内存
	a = new(int) // new 传入一个类型,返回该类型的指针
	*a = 10 // 可以正常赋值了
	fmt.Println(a) // 指针地址
	fmt.Printf("%T\n", a)
	var b int = 10
	b = 20
	fmt.Println(b)
}

func pointerTest()  {
	var a int = 10
	var b *int = new(int)
	b = &a
	*b = 20
	fmt.Println(a)
}

func pointerStruct()  {
	p := new(person)
	p.name = "ad"
	p.age = 10
	fmt.Println(p)
	// 相当于new了一遍person
	b := &person{}
	b.name = "b" // 这一步 = (*b).name 是go语言帮我们实现的语法糖
	b.age = 1
}

func interview()  {
	a := map[string]*person{}
	list := []person{
		{name: "ad", age: 1},
		{name: "ad", age: 2},
		{name: "ad", age: 3},
	}
	for _, val := range list {
		a[val.name] = &val
	}
	for key,val := range a {
		fmt.Println(key, "=>", val)
	}
}

func main()  {
	interview()
}