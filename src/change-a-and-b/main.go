package main

import "fmt"

/**
 * @ Author     ：zhaolengquan.
 * @ Date       ：Created in 22:43 2022/4/15
 * @ Description：交换a,b的值
 */
func main() {
	a := 3
	b := 4
	a, b = b, a
	fmt.Println(a)
	fmt.Println(b)
	//第二种方案
	println(change2(10, 20))

}
