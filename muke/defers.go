package main

import "fmt"

func f1() int {
	i := 5
	defer func() {
		i++
	}()
	return i
}

func f2(result int) int {
	defer func() {
		result++
	}()
	return 0
}

func f3() int {
	i := 5
	defer func() {
		i = i + 5
	}()
	return i
}

func f4(r int) int {
	defer func(r int) {
		r = r + 5
	}(r)
	return r
}

func f5Add(x, y int) (z int) {
	// 通过defer滞后性累加 -> 结果为100 + x + y
	defer func() {
		z = z + 100
	}()
	// 具名返回值,返回
	z = x + y
	return z
}

func sliceDefer() {
	var s [5]struct{}
	for i := range s {
		// 4,4,4,4 类似js异步
		defer func() { fmt.Println(i) }()
	}
}

// defer 代码块会在函数调用链表中增加一个函数调用. 这个函数调用不是普通的函数调用，而是会在函数正常返回，也就是return之后添加一个函数调用。因此，defer通常用来释放函数内部变量

func main()  {
	fmt.Println(f1())
	fmt.Println(f2(1))
	fmt.Println(f3())
	fmt.Println(f4(3))
	sliceDefer()
}
