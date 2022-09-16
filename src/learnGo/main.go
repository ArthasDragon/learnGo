package main

import "fmt"

var age int

var level = 1
var (
	a = 1
	b = 2
)

func main() {
	//s3 := "luanran"
	//level = 2
	//fmt.Println(s3)
	//fmt.Printf("%T,%d", level, level)

	//	net
	//var conn net.Conn
	//var err error

	//conn, err := net.Dial("tcp", "www.baidu.com")
	//conn, _ := net.Dial("tcp", "www.baidu.com")
	//fmt.Println(conn)
	//
	//a, b = b, a
	//fmt.Println(a, b)

	//var a = 100
	//var b = 200
	//c := sum(a, b)
	//fmt.Printf("sum=%d", c)

	//	切片
	//var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//arr1 := arr[0:4]
	//arr2 := arr[2:6]
	//fmt.Println(len(arr1), cap(arr1))
	//fmt.Println(len(arr2), cap(arr2))
	//
	//arr3 := arr2[2:4]
	//fmt.Println(len(arr3), cap(arr3), arr3)

	//s1 := make([]int, 3, 10)
	//fmt.Printf("s1：%v len(s1)：%d cap(s1)：%d\n", s1, len(s1), cap(s1))
	//s1 = []int{1, 2, 3}
	//fmt.Printf("s1：%v len(s1)：%d cap(s1)：%d\n", s1, len(s1), cap(s1))
	//
	//for i := 0; i < len(s1); i++ {
	//	fmt.Println(s1[i])
	//}
	//for i, v := range s1 {
	//	fmt.Println(i, v)
	//}
	//s2 := []int{4, 5}
	//s1 = append(s1, s2...)
	//s3 := make([]int, 6)
	//copy(s3, s1)
	//fmt.Printf("s1：%v len(s1)：%d cap(s1)：%d\n", s1, len(s1), cap(s1))
	//fmt.Println(s3)
	//
	////	删除切片中的元素
	//s1 = append(s1[:1], s1[2:]...)
	//fmt.Println(s1)

	////	数组转切片
	//var s4 = [...]int{1, 2, 3}
	//fmt.Printf("s4类型：%T\n", s4)
	//s5 := s4[:]
	//fmt.Printf("s5类型：%T\n", s5)
	//
	////	数组
	//var all [2][3]int
	//all = [2][3]int{
	//	[3]int{1, 2, 3},
	//	[3]int{4, 5, 6},
	//}
	//
	//fmt.Println(all)

	// map
	//var m1 map[string]int
	//fmt.Println(m1)
	//
	//var m2 map[string]int
	//m2 = make(map[string]int, 10)
	//m2["lucky"] = 18
	//m2["jason"] = 24
	//fmt.Println(m2, m2["lucky"], m2["12323"])
	//
	//value, ok := m2["lucky"]
	//fmt.Println(value, ok)
	//
	//for k, v := range m2 {
	//	fmt.Println(k, v)
	//}

	////	rune
	//s := "hello王"
	//sHello := "hello"
	//sWang := "王"
	//fmt.Println(len(s))
	//fmt.Println(len(sHello))
	//fmt.Println(len(sWang))
	//
	//for _, c := range s {
	//	fmt.Printf("%c\n", c)
	//}
	//
	//// 字符串的修改要转成rune切片，而不能像PHP一样直接修改
	//s1 := []rune(s)
	//s1[0] = 'd'
	//fmt.Println(string(s1))
	//
	//c1 := "d"
	//c2 := 'd'
	//fmt.Printf("c1的类型：%T c2的类型：%T \n", c1, c2)
	////	总结：我们发现只要是双引号包裹的类型就是string，只要是单引号包裹的类型就是int32，也就是rune。和中英文无关。

	////	pointer
	////  go语言不存在指针操作，只有2个符号： &取内存地址 *根据内存地址取值
	//n := 18
	//fmt.Println(&n)
	//fmt.Println(*&n)
	//
	////  查看内存地址的类型
	//p := &n
	//fmt.Printf("%T\n", p) //  结果是*int  int类型的指针
	////	总结
	////  对变量进行取地址操作（&），可以获得这个变量的指针变量
	////  指针变量的值是指针地址（内存地址）
	////  对指针变量进行取值操作（*），可以获得这个指针变量指向原变量的值，即通过内存地址取值。

	//	流程控制
	// 	for i := 0; i < 10; i++ {
	// 		fmt.Println(i)
	// 	}
	// 	s := "hello"
	// 	for i, i2 := range s {
	// 		fmt.Printf("ikey值为：%v, i2字符为：%c \n", i, i2)
	// 	}

	// 	//  goto   不推荐使用
	// 	for i := 0; i < 10; i++ {
	// 		for j := 'A'; j < 'Z'; j++ {
	// 			if j == 'C' {
	// 				goto ccc
	// 			}
	// 			fmt.Printf("%v-%c \n", i, j)
	// 		}
	// 	}
	// ccc:
	// 	fmt.Println("over")

	// 函数  闭包
	f1 := adder(1)
	ret := f1(2)
	fmt.Println(ret)

}

func adder(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}
func sum(a, b int) int {
	return a + b
}
