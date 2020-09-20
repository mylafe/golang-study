## 切片

Go 语言切片是对数组的抽象。

切片是一种动态数组，比数组操作灵活，长度不是固定的，可以进行追加和删除。

`len()` 和 `cap()` 返回结果可相同和不同。

#### 声明

    var 切片名 []类型
    比如：var a []int

```go
package main

import "fmt"

func main() {
	//定义一个数组
	var intArr [5]int = [...]int{1, 22, 33, 66, 99}
	//定义一个切片
	//slice是一个切片
	//引用了intArr数组的，起始下标为1，终止下标为3(但是不包含3)
	slice := intArr[1:3]
	fmt.Println("intArr数组元素是 ", intArr)
	fmt.Println("slice切片的元素是 ", slice)

	fmt.Println("slice切片的长度是 ", len(slice))
	fmt.Println("slice切片的容量是 ", cap(slice)) //容量 cap
	//打印intArr[1]的内存中的地址  %p 格式化为0x开头的十六进制
	fmt.Printf("intArr[1]的地址=%p\n", &intArr[1])
	//打印slice的地址
	fmt.Printf("slice的地址=%p\n", &slice)
	slice[1] = 34
	fmt.Println()
	fmt.Println()

	fmt.Println("intArr数组元素是 ", intArr)
	fmt.Println("slice切片的元素是 ", slice)
}
```



## 

