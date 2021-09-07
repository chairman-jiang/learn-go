package main

import (
	"fmt"
	"reflect"
	"strconv"
	"unsafe"
)

// "github.com/gin-gonic/gin"

func test() {
	// var a int = 10

	// fmt.Printf("值为%v，类型为%T", a, a)
	// b := float64(a)
	// fmt.Printf("b：%.2f%T", b, b)

	// var str = `qwe-
	// asd-
	// zxc`
	// fmt.Println(str)
	// var s = strings.Split(str, "-")
	// fmt.Println(s)

	// var str2 = "wqeqweqwe01232130"
	// var ss = fmt.Sprintf("%v\n%v", str, str2)
	// fmt.Println(ss)
	// fmt.Println(strings.Split(str2, "0"))
	// var uns = strings.Join(s, "+")
	// fmt.Println(uns)
	// var p = strings.Split(str, "\n")
	// fmt.Println(p)

	// var str3 = "qwe"
	// var syb = strings.HasPrefix(str3, "q")
	// var sbh = strings.HasSuffix(str3, "e")
	// var idx = strings.Index(str3, "w")
	// var lastIdx = strings.LastIndex(str3, "e")
	// fmt.Printf("%v%v%v%v", syb, sbh, idx, lastIdx)
	// ---------------------------------------------
	// var a = "aasd"
	// for i := 0; i < len(a); i++ {

	// 	fmt.Printf("%v-%c\n", a[i], a[i]) // 在遍历汉字或其他utf8字符时，用range关键字
	// }
	// var b = "阿什杜"
	// for index, item := range b {
	// 	fmt.Printf("item：%vindex:%v-val：%c", item, index, item)
	// }

	// 想要修改字符串中某个字符需要转换类型，字符串中存在utf8字符集需要使用rune处理
	// var c = "hello" // ascll 型
	// var c1 = []byte(c)
	// c1[0] = 'a'
	// fmt.Println(c1)

	// var d = "阿什杜" // utf-8 型
	// var d1 = []rune(d)
	// d1[0] = 'a'
	// fmt.Println(d1)

	// go字符集分为两种 byte（ASCLL码一个字符） 和 rune（代表一个utf-8字符）
	// 将其他简单数据类型转换为string类型时可以使用fmt.Sprintf()
	// 其中占位符需要注意 int -> %d  float -> %f  bool -> %t  byte -> %c
	// 还可以使用strconv这个包 进行 字符串转换

	// 使用除法注意，两位都是整数，结果也是整数
	// 10 / 3 = 3
	// 10.0 / 3.0 = 3.33333333
	// 取余 -> 余数 = 被除数-（被除数/除数）*除数
	// -10 % 3  -> -10 - (-10/3)*3 = -1

	// golang中 ++ -- 都是单独的语句并不是运算符
	// golang中 不能再赋值操作后面使用， 不能使用前置++ --
	// var e int = 1;
	// var f int = ++e; 这是错误的写法

	var tmp int = 1
	for i := 1; i <= 10; i++ {
		tmp *= i
	}
	fmt.Printf("%v", tmp)
}

func testArr() {
	var arr = []int{1, 3}
	fmt.Printf("%v", arr)
}

func testMap() {
	var map1 = make([]map[string]string, 3, 3)
	fmt.Printf("%v", map1)
}

func swap(a, b int) {
	a, b = b, a
	fmt.Println("swap => a:", a, "g:", b)
	//var mode = flag.String("mode", "", "process mode")
	//	str := new(string)
}

/*
	a,b 类型都是指针类型
	&a: 取变量a的地址, 传入形参
	函数形参传入的永远只是值
*/
func swap2(a, b *int) {
	*a, *b = *b, *a
	fmt.Println("a:", a, "b", b)
}

/*
	切片:
		数组的容量固定, 不能自动扩展
		值传递, 数组作为函数传递参数时, 将整个数组的值拷贝一份给形参
*/

func sliceTest() {
	arr := []int{1, 2, 3, 4, 5, 6}
	// [low : high : max] 最低位, 最高位, 容量
	// 如果截取时没有指定max, 那么容量就就是原数组的的容量
	s := arr[1:3:5]

	a := arr[1:]

	// len: 长度  cap: 容量
	fmt.Println("S: ", s, "A: ", a)

	// 切片使用 append 时, 如果容量不足, 1024以下 会以当前两倍的容量进行扩容, 1024以上修改扩容算法 但是切片的内存地址会发生改变
}

func clearEmpty(list []string) []string {
	//var data = make([]string, 0)
	//for i, t := range list {
	//	if t != "" {
	//		fmt.Println("item => ", i, t)
	//		data = append(data, t)
	//	}
	//}
	//return  data
	i := 0 // 记录着切片的最后一个下标
	for _, t := range list {
		if t != "" {
			list[i] = t
			i++
		}
	}
	return list[0:i]
}

func clearSame(list []string) []string {
	//newList := list[:1]
	//for _, t := range list {
	//	isHas := false
	//	for _, sub := range newList {
	//		if t == sub {
	//			isHas = true
	//			break
	//		}
	//	}
	//	if !isHas {
	//		newList = append(newList, t)
	//	}
	//}
	//return newList

	newList := list[:1]
	for _, t := range list {
		i := 0
		for ; i < len(newList); i++ {
			if t == newList[i] {
				break
			}
		}
		fmt.Println("i", i)
		if i == len(newList) {
			newList = append(newList, t)
		}
	}
	return newList
}

func copyTest() {
	arr := [...]int{1, 2, 3, 4, 2, 4}

	s := arr[2:3]

	c := arr[3:4]

	copy(s, c)
	fmt.Println("s:", s, "c", c)
}

func stringTest() {
	str := "你好 star"
	// 这样遍历是取ascll
	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}
}

func mapTest() {
	// map 自动扩容
	// key 不能是引用类型
	// obj := make(map[int]string, 5)
	var map1 map[int]string = map[int]string{1: "123"}
	map1[11] = "123"
	for key, val := range map1 {
		fmt.Println("key", key, "val", val)
	}
	// 对map类型的数据取值
	v, has := map1[22]
	fmt.Println("v", v, "has", has)
	delete(map1, 1)
}

type Person struct {
	name string
	age  int
}

func testStruct(myType Person) {
	/*
		结构体比较智能 == 和 !=
		相同结构体(成员变量的类型, 个数, 顺序 要一致) 变量之间是可以直接赋值的
		结构体函数传参, 值传递 内存消耗大 效率低
		结构体函数传参使用地址传递
		地址: 结构体元素的地址 == 结构体首个元素的地址
	*/
	/*
		创建结构体指针变量
	*/
	//var man *Person = &Person{"asd", 23}
	fmt.Println("test -> sizeOf", unsafe.Sizeof(myType))
}

func testGoto() {
	i := 1
flag:
	if i <= 5 {
		fmt.Println(i)
		i++
		goto flag
	}

	j := 5
	for {
		if j >= 10 {
			goto flag2
		}
		j++
		fmt.Println(j)
	}

flag2:
}

type Profile struct {
	name   string
	age    int
	mother *Profile
	father *Profile
}

// 绑定到Profile结构体 以值作为方法接收者
func (person Profile) FmtProfile() {
	fmt.Println(person.name)
	fmt.Println(person.age)
}

// 加上 * 可改变实例内部值 以指针作为方法接收者
func (person *Profile) ChangeProperty() {
	person.age += 1
}

type company struct {
	companyName string
}

type people struct {
	name    string
	company // 匿名字段 就是要被继承的结构体的名称(company)
}

/*
	myCompany := company{companyName: "视联动力"}
	// company: myCompany 实现继承
	me := people{ name: "chris", company: myCompany }
*/

/*
	当方法的首字母为大写时，这个方法对于所有包都是Public，其他包可以随意调用
	当方法的首字母为小写时，这个方法是Private，其他包是无法访问的
*/

type Good interface {
	sumToGoodPrice() int
	orderInfo() string
}

type phone struct {
	name     string
	quantity int
	price    int
}

func (myphone phone) sumToGoodPrice() int {
	return myphone.price * myphone.quantity
}

func (myphone phone) orderInfo() string {
	return "您要购买" + myphone.name + strconv.Itoa(myphone.quantity) + "件. 共计" + strconv.Itoa(myphone.sumToGoodPrice()) + "元"
}

type freegood struct {
	name     string
	quantity int
	price    int
}

func (myGoods freegood) sumToGoodPrice() int {
	return myGoods.price * myGoods.quantity
}

func (myGoods freegood) orderInfo() string {
	return "您要购买" + myGoods.name + strconv.Itoa(myGoods.quantity) + "件. 共计" + strconv.Itoa(myGoods.sumToGoodPrice()) + "元"
}

func totalConsumption(goods []Good) int {
	var count int
	for _, item := range goods {
		fmt.Println(item.orderInfo())
		count += item.sumToGoodPrice()
	}
	return count
}

/*
payPhone := phone{name: "iphone", price: 5000}
payFreegood := freegood{name: "toys", price: 300}

goods := []Good{payPhone, payFreegood}

price := totalConsumption(goods)

fmt.Println(price)

*/

type Person2 struct {
	Name string `label:"Name is:" default:"123"`
}

func PrintTest(obj interface{}) {
	fmt.Println(reflect.TypeOf(obj))
	v := reflect.ValueOf(obj)
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		tag := field.Tag
		label := tag.Get("label")
		defaultValue := tag.Get("default")
		fmt.Println(label, defaultValue)
		fmt.Println(v.Field(i))
	}
}

func er(num int)  {
	str := strconv.Itoa(num)
	fmt.Println(len(str))
	for i:= 0; i < len(str); i++ {
		fmt.Println(str[i])
	}
	for index, item := range str {
		fmt.Println(index, item)
	}
}

func main() {
	er(123)
	//p2 := Person2{Name: "XuChen"}
	//PrintTest(p2)
	// var per Person = Person{"123", 22}
	// fmt.Println("main -> sizeOf", unsafe.Sizeof(per))
	// testStruct(per)
	//var i int = 1
	//fmt.Println("i", &i)
	//var ptr *int = &i
	//fmt.Printf("%v\n", ptr)
	//*ptr = 12
	//fmt.Println("ptr",ptr)

	//var a int = 1
	//var b int = 3
	//a, b = b, a
	//a = a + b
	//b = a - b
	//a = a - b
	//fmt.Println(a,b)
	//var (
	//	name string
	//	age int
	//)
	//fmt.Println("请输入姓名")
	//n, err := fmt.Scanln(&name)
	//if err != nil {
	//	panic(err)
	//} else {
	//	fmt.Println(n)
	//}
	//
	//fmt.Println("请输入年龄")
	//fmt.Scanln(&age)
	//fmt.Printf("请检查输入是否正确, 姓名%s-年龄%d", name, age)
}
