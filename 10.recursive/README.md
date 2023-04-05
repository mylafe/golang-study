## 递归

Go 语言支持递归。但我们在使用递归时，开发者需要设置退出条件，否则递归将陷入无限循环中。

递归函数对于解决数学上的问题是非常有用的，就像计算阶乘，生成斐波那契数列等。

格式：

    func recursion() {
       recursion() /* 函数调用自身 */
    }
    
    func main() {
       recursion()
    }

#### 实例
- 阶乘

```go
package main

import "fmt"

func Factorial(n uint64)(result uint64) {
    if n > 0 {
        result = n * Factorial(n-1)
        return result
    }
    return 1
}

func main() {  
    var i int = 15
    fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))
}
```

> D:\Program Files (x86)\Ampps\www\github\golang-study\10.recursive>go run fact.go

    15 的阶乘是 1307674368000

- 斐波那契数列

```go
package main

import "fmt"

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func main() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}
}
```

> D:\Program Files (x86)\Ampps\www\github\golang-study\10.recursive>go run fib.go

    0       1       1       2       3       5       8       13      21      34

[首页](../README.md)

[下一章](../11.interface/README.md)
