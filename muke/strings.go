package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func strTest()  {
	str1 := "yes,we-are-the-world"
	for i, ch := range str1 {
		// string 进行utf-8 解码 再转为Unicode
		fmt.Printf("(%d, %X):", i, ch)
	}
	fmt.Println()
	/*
		RuneCountInString: 获取字符长度
		DecodeRune: 每次只解析一个(rune)
	*/
	runes := utf8.RuneCountInString(str1)
	fmt.Println(runes)
	bytes := []byte(str1)
	for len(bytes) > 0 {
		r, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c", r)
	}
	fmt.Println()

	for i, ch := range []rune(str1) {
		// 就可以准确拿到字符串下标
		fmt.Printf("(%d, %c)", i,ch)
	}
}

func test2() {
	// 字符串底层实际是一个byte数组
	var str1 string = "123"
	var str2 string = "456"
	// 拼接
	var str3 = strings.Join([]string{str1, str2}, "-")
	var str4 = strings.Split(str3, "-")
	fmt.Println(str3, str4)
	// rune类型是一个int32
	var str5 string = "严abc"
	for val, i := range []byte(str5) {
		fmt.Println(val, i)
	}
	fmt.Println(len(str5))
	fmt.Println(utf8.RuneCountInString(str5))
}

func transformation() {
	// unicode 转换 utf8 规则 具体参见 https://www.ruanyifeng.com/blog/2007/10/ascii_unicode_and_utf-8.html
	/*
	unicode是什么
		unicode又被称为万国码,是一种编码集
		由于计算机只能识别二进制,所有需要一套字符集映射字符 {二进制: 字符}
		unicode采用16进制key映射value, 但是ASCII码都是些英文是需要1个字节存储, unicode会造成内存浪费
	由uft-8来解决这些问题

	unicode为16进制
	1.现将16进制转为2进制
	2.寻找二进制中(1)的最大子串
	3.得出需要几个字节存储

	*/
}

func main()  {
	test2()
}
