package main

import "fmt"

/**
 * @ Author     ：zhaolengquan.
 * @ Date       ：Created in 08:51 2022/4/18
 * @ Description：
 */
func main() {

	//a := 10
	//if a > 10 {
	//	fmt.Println("a >10")
	//} else {
	//	fmt.Println("a<=10")
	//}
	//
	//if b := 10; b > 10 {
	//	fmt.Println("b>10")
	//} else if b < 10 {
	//	fmt.Println("b<10")
	//} else {
	//	fmt.Println("b==10")
	//}
	//
	//sum := 0
	//for i := 0; i < 100; i++ {
	//	if i%2 != 0 {
	//		continue
	//	}
	//	sum += i
	//}
	//fmt.Println("sum的值是", sum)

	//数组的定义,数组长度可以忽略, 也可以指定index进行初始化.
	//array := []string{1: "index1", 2: "index2", 3: "index3"}
	//for i := 0; i < len(array); i++ {
	//	fmt.Printf("数组索引是%d,数组值为%s\n", i, array[i])
	//}

	//for range新型循环遍历数组
	//i也可用_匿名变量代替(丢弃)
	//for i, s := range array {
	//	fmt.Println(i, s)
	//}

	//切片,array[i:j] 包含左边,不包含右边
	//i和j也能省略, 如果省略i,默认是start, 省略j 默认就是数组长度. 都省略 就是整个数组
	array := []int{1, 2, 3, 4, 5}
	slice := array[2:5]
	fmt.Println(slice)
	fmt.Println(array[:])
	fmt.Println(array[2:])
	fmt.Println(array[:3])

	//切片的值允许被修改,也就证实了切片也是数组
	slice = array[2:4]
	fmt.Println(slice)
	slice[0] = 1
	fmt.Println(slice)

}
