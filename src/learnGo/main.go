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
	var m1 map[string]int
	fmt.Println(m1)

	var m2 map[string]int
	m2 = make(map[string]int, 10)
	m2["lucky"] = 18
	m2["jason"] = 24
	fmt.Println(m2, m2["lucky"], m2["12323"])

	value, ok := m2["lucky"]
	fmt.Println(value, ok)

	for k, v := range m2 {
		fmt.Println(k, v)
	}

}

func sum(a, b int) int {
	return a + b
}
