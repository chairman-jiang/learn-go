package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

/*
	slice 本身是没有数据的, 是对底层array的一个view
	slice 可以向后扩展, 不可以向前
	slice 有三部分组成 ptr len cap(capacity)
		ptr: 指向slice开头的那个元素
		len: 本身长度 --- 可通过len()函数获取长度
		cap: 底层view的array的长度 --- 可通过cap()函数获取长度
*/


func printArray() {
	var arr = [...]int{1,2,3,4,5,6,7,8,9,10}

	arr1 := arr[:]

	arr1[1] = 100
	fmt.Println(arr1)
	/*
		切片 arr1 下标0位置变为100
		原数组 arr 下标0位置也变为100
	*/
}

func add() {
	var arr = [...]int{1,2,3,4,5,6,7,8}
	s1 := arr[5:7]
	s2 := append(s1, 8)
	s3 := append(s2, 9)
	s4 := append(s3, 10)
	// 以上 slice no longer view array
	// 添加元素时,超过cap, 系统会重新分配更大的底层数组
	fmt.Println("s4:", s4)
	fmt.Println("arr:", arr)

	// 由于按值传递的关系, append必须有接收值
}

func printSlice(s []int) {
	fmt.Printf("len=%d,cap=%d", len(s), cap(s))
}

func loopSlice()  {
	var s []int
	/*
		变量定义有个 zero value (初始化的值)
		Zero value for slice is nil
	*/
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, i)
	}
	fmt.Println("s: ", s)
}

func sum(arr [10]int, target int) (firstIndex int, lastIndex int) {
	var first int = -1
	var last int = -1
	for i := 0; i < len(arr); i++ {
		if first > -1 {
			break
		}
		for j := i; j < len(arr); j++ {
			if arr[i] + arr[j] == target {
				first = i
				last = j
				break
			}
		}
	}
	return first, last
}
/*
	slice自动扩容
*/
func autoDistribution()  {
	a := make([]int, 1, 2)
	s := cap(a)
	for i := 0; i < 50; i++ {
		a = append(a, i)
		if n := cap(a); n > s {
			fmt.Printf("cap -> %d -> %d\n", s, n)
			s = n
		}
	}
}

/*
	copy注意事项
*/
func copied() {
	/*
		copy函数是对两个slice间数据的复制, 两个slice操作同一个底层数组, 在copy时, 底层源数组改变, 对应的slice内部的值也改变
	*/
	s := [5]int{1,2,3,4,5}
	// 2,3,3,4,5
	s1 := s[0 : 2]
	s2 := s[1 : 3]
	fmt.Println(s1, s2)
	copy(s1, s2)
	fmt.Println(s1, s2, s)
}

/*
	slice resize
*/
func resizeSlice() {
	s := []int{1,2,3,4}
	b := s[1 : 2] // b = 2; 此时 b的ptr就是 2,3,4 因为从下标1开始截取的
	// high位最大为 b的ptr的len, 否则保存
	c := b[0 : 3] // b的len只是1, 超出就去b的ptr取值 所以就是 2,3,4
	fmt.Println(b, c)
}

func stringSlice() {
	str := "你好，世界！hello world！"
	strList := []rune(str)
	strList[3] = '狗'
	strList[4] = '浪'
	strList[12] = 'g'
	strList = strList[:14]
	str = string(strList)
	fmt.Println(str)
}

func arrayOrSliceConvertToString()  {
	arr := [10]int{1,2,3,4,5,6,7,8,9,10}
	slice := make([]int, 10)
	//rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		slice[i] = rand.Intn(100)
	}
	fmt.Println(slice)
	str3 := strings.Replace(strings.Trim(fmt.Sprint(arr),"[]")," ", ",", -1)
	fmt.Println(str3)
}

type Student struct {
	name string
	age int
}

func change(m map[string]Student)  {
	m["first"] = Student{name: "456", age: 2}
}

func stuTest()  {
	s := map[string]Student{}
	s["first"] = Student{name: "123", age: 1}
	fmt.Println(s, "-> first")
	change(s)
	fmt.Println(s, "-> second®")
}

func mySort(m map[string]Student) []Student {
	list := make([]int, 0, 10)
	copiedStu := make([]Student, 0, 10)
	for _, val := range m {
		list = append(list, val.age)
	}
	sort.Ints(list)
	for i := 0; i < len(list); i++ {
		for _, val := range m {
			if val.age == list[i] {
				copiedStu = append(copiedStu, val)
			}
		}
	}
	return copiedStu
}

func main() {
	stuMap := map[string]Student{}
	stuMap["first"] = Student{name: "Kris", age: 1}
	stuMap["second"] = Student{name: "Mike", age: 2}
	stuMap["third"] = Student{name: "Soul", age: 5}
	stuMap["fourth"] = Student{name: "Christian", age: 4}
	stuMap["fifth"] = Student{name: "Neil", age: 9}
	copiedStu := mySort(stuMap)
	for _, val := range copiedStu {
		fmt.Println(val)
	}
}
