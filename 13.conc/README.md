## 并发

goroutine 是轻量级线程，goroutine 的调度是由 Golang 运行时进行管理的。

goroutine 语法格式：

    go 函数名( 参数列表 )

#### 实例

```go
package main

import (
        "fmt"
        "time"
)

func say(s string) {
        for i := 0; i < 5; i++ {
                time.Sleep(100 * time.Millisecond)
                fmt.Println(s)
        }
}

func main() {
        go say("world")
        say("hello")
}
```

输出的 hello 和 world 是没有固定先后顺序。因为它们是两个 goroutine 在执行

> D:\Program Files (x86)\Ampps\www\github\golang-study\13.conc>go run demo.go

    world
    hello
    world
    hello
    hello
    world
    world
    hello
    world
    hello

#### 通道

通道（channel）是用来传递数据的一个数据结构。

通道可用于两个 goroutine 之间通过传递一个指定类型的值来同步运行和通讯。操作符 <- 用于指定通道的方向，发送或接收。如果未指定方向，则为双向通道。

```go
ch <- v    // 把 v 发送到通道 ch
v := <-ch  // 从 ch 接收数据
           // 并把值赋给 v
```

声明一个通道很简单，我们使用chan关键字即可，通道在使用前必须先创建：
    
    ch := make(chan int) //通道可以设置缓冲区，通过 make 的第二个参数指定缓冲区大小
    
    
###### 实例

```go
package main
 
 import "fmt"
 
 func sum(s []int, c chan int) {
         sum := 0
         for _, v := range s {
                 sum += v
         }
         c <- sum // 把 sum 发送到通道 c
 }
 
 func main() {
         s := []int{7, 2, 8, -9, 4, 0}
 
         c := make(chan int)
         go sum(s[:len(s)/2], c)
         go sum(s[len(s)/2:], c)
         x, y := <-c, <-c // 从通道 c 中接收
 
         fmt.Println(x, y, x+y)
 }
```
> -5 17 12

#### Go 遍历通道与关闭通道

Go 通过 range 关键字来实现遍历读取到的数据，类似于与数组或切片。格式如下:

    v, ok := <-ch
 
###### 实例

```go
package main

import (
        "fmt"
)

func fibonacci(n int, c chan int) {
        x, y := 0, 1
        for i := 0; i < n; i++ {
                c <- x
                x, y = y, x+y
        }
        close(c)
}

func main() {
        c := make(chan int, 10)
        go fibonacci(cap(c), c)
        // range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
        // 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
        // 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
        // 会结束，从而在接收第 11 个数据的时候就阻塞了。
        for i := range c {
                fmt.Println(i)
        }
}
```

> D:\Program Files (x86)\Ampps\www\github\golang-study\13.conc>go run goRun.go

    0
    1
    1
    2
    3
    5
    8
    13
    21
    34
