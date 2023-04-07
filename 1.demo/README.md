#### demo

#### 代码

````go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

````

#### 执行结果

> D:\Program Files (x86)\Ampps\www\github\golang-study\1.demo>go run helloworld.go
Hello, World!

#### 记一次踩坑

go run 报错
> textflag.h:1: '#' must be first item on line

go env都正常。折腾了一段时间，打开对应文件发现textflag.h乱码了，~~亿赛通加密软件问题~~。

    原因：头文件加密导致go.exe无法正确读取textflag.h
    修复：杀进程——修改明文模式——重装go环境——go run success！
