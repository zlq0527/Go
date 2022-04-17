package main

/**
 * @ Author     ：zhaolengquan.
 * @ Date       ：Created in 23:22 2022/4/15
 * @ Description：
 */
var name1 = "abc"

func main() {
	//局部变量
	//var test = 3
	//fmt.Println(test)
	//a := 10
	//b := 20
	//c := a + b
	//fmt.Printf("a=%d,b=%d,c=%d", a, b, c)

	//全局变量
	//fmt.Println(name1)
	//var name1 = 10
	//fmt.Println(name1)

	println(test(20, 3))
}
func test(a, b int) (int, int) {
	return b, a
}
