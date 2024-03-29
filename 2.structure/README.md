#### 语言结构

#### 结构

- 包声明
- 引入包
- 函数
- 变量
- 语句 & 表达式
- 注释

#### 实例

helloworld.go为例

````go
package main

import "fmt"

func main() {
    /* hello world */
    fmt.Println("Hello, World!")
}
````

> package main 定义了包名
>
> import "fmt"   import 导入依赖的包
>
> func main()是程序入口
>
> /* 注释内容 */ 程序注释
>
> fmt.Println("Hello, World!") 一个输出语句，使用 fmt.Print("hello, world\n") 可以得到相同的结果。
>
> 当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，如：Group1，那么使用这种形式的标识符的对象就可以被外部包的代码所使用（客户端程序需要先导入这个包），这被称为导出（像面向对象语言中的 public）；标识符如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可见并且可用的（像面向对象语言中的 protected ）。

    总结：
    
    只有package 名为main的包才可以包含main()函数
    
    一个可执行程序有且只有一个main()函数
    
    通过import 可以导入其他依赖包（非main包），可以一次导入多个包
    
#### 规范

1. 注释，单行注释 // 开头，后面跟注释内容

2. 多行注释 /**/    * 号中间写注释内容，有其他语言基础的同学一看就能懂。

3. 变量生命必须要用空格 var name type

4. go语言的结尾不需要分号结尾，换行就代表一个语句结束。

5. 可见性规则，使用大小写来决定是否可以被外部包所调用。大写字母开头可以被外部代码调用，小写字母开头，外部代码不可调用，相当于面向对象语言中的public和private

6. `shift + alt + F` 格式化代码（vscode），或者使用 go fmt
    
#### 注意

func main() { 程序入口 `{` 不能单独放在一行，会产生错误

[首页](../README.md)

[下一章](../3.dataType/README.md)
