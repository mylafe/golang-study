package main

import "fmt"

func main() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200
	var ret int

	/* 调用函数并返回最大值 */
	ret = max(a, b)

	fmt.Printf( "最大值是 : %d\n", ret )
}

/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
	/* 定义局部变量 */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}
