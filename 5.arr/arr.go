package main

import (
	"fmt"
)

func main() {
	//一维数组
	var arr1 [5] int
	fmt.Println(arr1)

	var arr2 = [5] int {1,2,3,4,5}
	fmt.Println(arr2)

	arr3 := [5] int {1, 2, 3, 4, 5}
	fmt.Println(arr3)

	arr4 := [...] int {1, 2, 3, 4, 5, 6}
	fmt.Println(arr4)

	arr5 := [5] int {0: 3, 1:5, 4:6}
	fmt.Println(arr5)

	//二维数组
	var arr6 = [3][5]int {{1, 2, 3, 4, 5},{9, 8, 7, 6, 5},{3, 4, 5, 6, 7}}
	fmt.Println(arr6)

	arr7 := [3][5]int {{1, 2, 3, 4, 5},{9, 8, 7, 6, 5},{3, 4, 5, 6, 7}}
	fmt.Println(arr7)

	arr8 := [...][5] int {{1, 2, 3, 4, 5},{9, 8, 7, 6, 5},{0: 3, 1:5, 4:6}}
	fmt.Println(arr8)
}
