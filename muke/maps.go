package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func mapTest()  {
	//var map1 map[string]string
	//var map2 = make(map[string]string)
	// key 是无序的, hash - name
	map3 := map[string]string {
		"name": "chris",
		"age": "12",
	}
	for i, n := range map3 {
		fmt.Println(i,n)
	}
	// 当取map中不存在的值时, 还是可以返回一个空串的(zero value)
	// 对map取值判断是否存在
	if val1, ok := map3["class"]; ok {
		fmt.Println("val1:", val1)
	} else {
		fmt.Println("key dose not exist :", val1)
	}

	value, ok := map3["cla"]
	fmt.Println(value, ok)
	value2, ok := map3["age"]
	fmt.Println(value2, ok)
}
// 无重复字符的最长子串
func lengthOfRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s){
		lastI, ok := lastOccurred[ch]
		if ok && lastI >= start {
			start = lastOccurred[ch] + 1
		}
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

// 循序遍历map
func order()  {
	// 开启随机数, 否则调用rand永远的同样的数
	rand.Seed(time.Now().UnixNano())
	m := make(map[string]int, 200)
	for i:= 0; i < 100; i++ {
		// sprintF是返回格式化后的字符串
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		m[key] = value
	}
	var keys = make([]string, 0, 200)
	for key, _ := range m {
		keys = append(keys, key)
	}
	// 按字符排序
	sort.Strings(keys)
	for _, val := range keys {
		fmt.Println(val)
	}
}

func mapSlice()  {
	m := make(map[string][]string, 10)
	m["one"] = []string{"123", "456"}
	fmt.Println(m)
}

func sliceMap()  {
	s := make([]map[string]int, 10)
	s[0] = map[string]int{}
	s[0]["age"] = 1
	fmt.Println(s)
}

func main()  {
	mapSlice()
	sliceMap()
}
