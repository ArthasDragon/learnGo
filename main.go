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
	//conn, err1 := net.Dial("tcp", "www.baidu.com")
	//fmt.Println(conn)
	//fmt.Println(err)
	//fmt.Println(err1)

	a, b = b, a
	fmt.Println(a, b)
}
