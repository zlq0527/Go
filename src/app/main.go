package main

import (
	"fmt"
)

func getdata() (int, float32, float64) {
	return 100, 133.224, 222.33333
}

func main() {

	a3, _, _ := getdata()
	_, _, a4 := getdata()
	fmt.Println(a3)
	fmt.Println(a4)

	//var float = 0.18
	//fmt.Println(float)
	//e.x = 10
	//e.b = "abcbac"
	//fmt.Println(e.b)
	//fmt.Println(e.x)
	//i, j := 0, "abc"
	//i = 20
	//fmt.Println("hello,world")
	//fmt.Println(greet.Name)
	//fmt.Println(i, j)

	//a := 349
	//b := 192
	//c := float32(a-b) * 0.1723244
	//fmt.Println(c)
	//var hp2 int
	//hp2 := 33
	//fmt.Println(hp2)

	//conn, err := net.Dial("tcp", "127.0.0.1:8080")
	//conn1, err := net.Dial("tcp", "127.0.0.1:8080")
	//fmt.Println(conn)
	//fmt.Println(err)
	//fmt.Println(conn1)

	//a := 1
	//b := 2
	//var t int
	//t = a
	//a = b
	//b = t
	//fmt.Println(a)
	//fmt.Println(b)

	//fmt.Println("----------")
	//c := 1
	//d := 2
	//c, d = d, c
	//fmt.Println(c)
	//fmt.Println(d)
}
