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

> D:\Program Files (x86)\Ampps\www\github\golang-study\9.slice>go run slice.go

    intArr数组元素是  [1 22 33 66 99]
    slice切片的元素是  [22 33]
    slice切片的长度是  2
    slice切片的容量是  4
    intArr[1]的地址=0xc0000102d8
    slice的地址=0xc0000044a0

#### 使用

1. 定义一个切片，用切片去引用前面创建好的数组

```go
var arr [5]int = [...]int{1, 2, 3, 4, 5}
slice := arr[1:3]
```

2. 通过make来创建切片

> 语法： var 切片名 [ ]type = make([ ]type , len , [cap])

> 参数说明：type 数据类型 len 长度 cap 容量

```go
func main() {
    var slice []int = make([]int, 4, 10)
    //默认值
    fmt.Println("默认值为：", slice)
    // 长度
    fmt.Printf("slice的长度为：%v\n", len(slice))
    // 容量
    fmt.Printf("slice的容量为：%v\n", cap(slice))
    slice[0] = 100
    slice[2] = 200
    fmt.Println(slice)
}
```

3. 定义一个切片，直接指定一个数组，使用原理类似make方式
```go
var slice []int = []int{1, 3, 5}
var strSlice []string = []string{"tom", "jack", "mary"}
```

#### 遍历

- for 和 for range两种方式

###### 实例

```go
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
```

> D:\Program Files (x86)\Ampps\www\github\golang-study\9.slice>go run sliceFor.go

    i=0 v=20 i=1 v=30 i=2 v=40
    i=0 v=20
    i=1 v=30
    i=2 v=40

#### 截取

###### 实例

```go
package main

import "fmt"

func main() {
	/* 创建切片 */
	numbers := []int{0,1,2,3,4,5,6,7,8}
	printSlice(numbers)

	/* 打印原始切片 */
	fmt.Println("numbers ==", numbers)

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])

	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:])

	numbers1 := make([]int,0,5)
	printSlice(numbers1)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2]
	printSlice(number2)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5]
	printSlice(number3)

}

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
```

> D:\Program Files (x86)\Ampps\www\github\golang-study\9.slice>go run sliceCut.go

    len=9 cap=9 slice=[0 1 2 3 4 5 6 7 8]
    numbers == [0 1 2 3 4 5 6 7 8]
    numbers[1:4] == [1 2 3]
    numbers[:3] == [0 1 2]
    numbers[4:] == [4 5 6 7 8]
    len=0 cap=5 slice=[]
    len=2 cap=9 slice=[0 1]
    len=3 cap=7 slice=[2 3 4]

[首页](../README.md)

[下一章](../10.recursive/README.md)
