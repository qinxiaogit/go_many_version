package main

import "fmt"
/**
 * 多个参数
 */
func sum(nums ...interface{}){
	for _,it := range nums{
		switch it.(type) {
		case int:
			fmt.Println("int:",it)
		case float64:
			fmt.Println("float:",it)
		case string:
			fmt.Println("string:",it)
		case []string:

			fmt.Println("string arr:",it)


		}
	}
	fmt.Println("test ini")
}
/**
 * 闭包
 */
func intSeq()func()int{
	i:=0
	return func() int {
		i++
		return i
	}
}
