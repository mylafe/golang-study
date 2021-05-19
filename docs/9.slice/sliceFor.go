package main

import "fmt"

func main() {
	var arr [5]int = [...]int{10, 20, 30, 40, 50}
	slice := arr[1:4] //20,30.40
	for i := 0; i < len(slice); i++ {
		fmt.Printf("i=%v v=%v ", i, slice[i])
	}
	fmt.Println()
	for i, v := range slice {
		fmt.Printf("i=%v v=%v \n", i, v)
	}
}
