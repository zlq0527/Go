package main

import (
	"strings"
)

/**
 * @ Author     ：zhaolengquan.
 * @ Date       ：Created in 21:54 2022/4/17
 * @ Description：
 */
func main() {
	//v1 := true
	//v2 := false
	//fmt.Println("v1=", v1, ",v2=", v2)
	//fmt.Println("v1=v2", v1 == v2)
	//var v3 int = 3
	//v4 := 4
	//fmt.Println("v3+v4=", v4+v3)
	//var s1 string = "hello"
	//var s2 string = " world"
	//var s3 string = (s1 + s2)
	//fmt.Println("s1的旧值", s1)
	//s1 = "aaa"
	//fmt.Println("s1的新值", s1)
	//s1 += s2
	//pi := &s1
	//fmt.Println("pi地址", pi)
	//fmt.Println("pi的值", *pi)
	//fmt.Println(s1)
	//fmt.Println(s3)

	//string1 := "112233"
	//atoi, err := strconv.Atoi(string1)
	//fmt.Println(string1, atoi, err)

	//var f1 float64 = 2.333
	//f2 := int(f1)
	//fmt.Println(f1, f2)

	//str := "hexo"
	//fmt.Println(strings.HasPrefix(str, "h"))
	//fmt.Println(strings.HasSuffix(str, "o"))
	println(strings.ContainsAny("abc", "ddd"))
	println(strings.Contains("abc", "ab"))

	println(strings.Compare("abc", "abc"))

}
