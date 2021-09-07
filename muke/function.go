package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"reflect"
	"runtime"
	"strconv"
)

const (
	b = 1 << (10 * iota)
	kb
	mb
	gb
	tb
	pb
)
type any interface {}
/**
 * 转二进制
 */
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}
/**
 * 读文件
 */
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// 一行
		fmt.Println(scanner.Text())
	}
}
/*
	两个返回值
*/
func div(a,b int) (int, int) {
	return a / b, a % b
}
/*
	返回值命名 外部调用时 编辑器自动命名变量已接收
*/
func div2(a,b int) (q, r int) {
	return a / b, a % b
}
/*
	也支持这样写 函数体过长时就不建议这样写
*/
func div3(a,b int) (q, r int) {
	q = a / b
	r = a % b
	return
}
/*
	golang 中一般两个返回值都是预期的值和错误对象搭配
*/
func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		// golang不使用的变量使用下划线 _ 声明, 否则保存
		return q, nil
	default:
		return 0, fmt.Errorf("unsupport operation: %s", op)
	}
}

func apply(op func(int, int) int, a,b int ) int {
	p := reflect.ValueOf(op).Pointer() // 获取函数指针
	opName := runtime.FuncForPC(p).Name()
	// 打印当前转入的func op 是谁 在哪一行
	fmt.Printf("Calling function %s with args " + "(%d, %d)\n", opName, a, b)
	return op(a, b)
}

// 1个或多个参数
func argsTest(a int, n ...int) int {
	var num int = a
	for _, v := range n {
		num += v
	}
	return num
}

// 1个字符串和多个任意类型的参数
func argsTestT(s string, list ...interface{})  {
	fmt.Println(s)
	fmt.Println(list)
}

func sumPrint(format string, n ...int) string {
	var num int = 0
	for _, v := range n {
		num += v
	}
	return fmt.Sprintf(format, num)
}

func twoResult() (one, two int) {
	return 1,4
}

func addTest(x, y int) int {
	return x + y
}

func sumTest(n ...int) int {
	var num int = 0
	for _, v := range n {
		num += v
	}
	return num
}

func forEach(list []map[string]int , fn func(insideVal interface{}, insideIndex int))  {
	for key, val := range list {
		fn(val, key)
	}
}

func sliceForMap(list []map[string]int, fn func(val int, index int) interface{}) []interface{} {
	var tempList []interface{}
	for insideKey, item := range list {
		tempList = append(tempList, fn(item["age"], insideKey))
	}
	return tempList
}

//func reduce(list []interface{}, fn func(prv, cur interface{}, index int, dataSource []interface{}) interface{}, target []interface{})  {}

// golang 只有值传递 slice map chan interface 在函数形参都是默认按引用传递

func main () {
	if result, err := eval(1, 4, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
	fmt.Println(apply(func(a,b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3,4))
	list := []int{1,2,3}
	// 这里的 slice 做形参需要展开
	sum := sumPrint("sum -> %d", list...)
	fmt.Println(sum)
	// 也可以使用函数返回值传递
	fmt.Println(addTest(twoResult()))
	fmt.Println(sumTest(twoResult()))

	var tempList []map[string]int = []map[string]int{
		{"age": 1},
		{"age": 2},
		{"age": 6},
	}
	forEach(tempList, func(insideVal interface{}, insideIndex int) {
		fmt.Println("val ->", insideVal, "index ->", insideIndex)
	})

	tempList2 := sliceForMap(tempList, func(val int, index int) interface{} {
		return map[string]int{
			"age": val * 2,
		}
	})
	fmt.Println(tempList2)
}
