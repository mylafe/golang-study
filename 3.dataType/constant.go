package main

import (
	"fmt"
)

func main() {
	const name string = "mylafe"
	fmt.Println(name)

	const age = 30
	fmt.Println(age)

	const name_1, name_2 string = "mylafe", "wave"
	fmt.Println(name_1, name_2)

	const name_3, age_3  = "mylafe", 30
	fmt.Println(name_3, age_3)
}
