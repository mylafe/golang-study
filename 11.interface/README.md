## 接口

Go 语言提供了另外一种数据类型即接口，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口

#### 定义

    /* 定义接口 */
    type interface_name interface {
       method_name1 [return_type]
       method_name2 [return_type]
       method_name3 [return_type]
       ...
       method_namen [return_type]
    }
    
    /* 定义结构体 */
    type struct_name struct {
       /* variables */
    }
    
    /* 实现接口方法 */
    func (struct_name_variable struct_name) method_name1() [return_type] {
       /* 方法实现 */
    }
    ...
    func (struct_name_variable struct_name) method_namen() [return_type] {
       /* 方法实现*/
    }

#### 实例

```go
package main

import (
	"fmt"
)

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func main() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

}
```

> D:\Program Files (x86)\Ampps\www\github\golang-study\11.interface>go run interface.go

    I am Nokia, I can call you!
    I am iPhone, I can call you!

[首页](../README.md)

[下一章](../12.error/README.md)
