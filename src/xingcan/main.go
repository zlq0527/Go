package main

import "fmt"

/**
 * @ Author     ：zhaolengquan.
 * @ Date       ：Created in 23:44 2022/4/15
 * @ Description：
	形式参数 函数名后面括号内的是形式参数,函数调用结束后就会被销毁,
	在函数未被调用时，函数的形参并不占用实际的存储单元，也没有实际值。
*/
func main() {
	var a = 10
	var b = 20
	fmt.Printf("main()函数中a=%d\n", a)
	fmt.Printf("main()函数中b=%d\n", b)
	c := sum(a, b)
	fmt.Printf("sum函数的返回值是%d", c)
}
func sum(a, b int) int {
	fmt.Printf("sum函数中a=%d\n", a)
	fmt.Printf("sum函数中b=%d\n", b)
	return a + b
}
