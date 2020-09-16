## 条件和循环

## 条件语句

#### if 语句

if 语句 由一个布尔表达式后紧跟一个或多个语句组成。

    if 布尔表达式 {
       /* 在布尔表达式为 true 时执行 */
    }

###### 实例

````go
package main

import "fmt"

func main() {
	/* 定义局部变量 */
	var a int = 10

	/* 使用 if 语句判断布尔表达式 */
	if a < 20 {
		/* 如果条件为 true 则执行以下语句 */
		fmt.Printf("a 小于 20\n" )
	}
	fmt.Printf("a 的值为 : %d\n", a)
}
````

> D:\Program Files (x86)\Ampps\www\github\golang-study\4.if-for>go run if.go

    a 小于 20
    a 的值为 : 10

#### if...else 语句

if 语句 后可以使用可选的 else 语句, else 语句中的表达式在布尔表达式为 false 时执行。

    if 布尔表达式 {
       /* 在布尔表达式为 true 时执行 */
    } else {
      /* 在布尔表达式为 false 时执行 */
    }
    
###### 实例

````go
package main

import "fmt"

func main() {
   /* 局部变量定义 */
   var a int = 100
 
   /* 判断布尔表达式 */
   if a < 20 {
       /* 如果条件为 true 则执行以下语句 */
       fmt.Printf("a 小于 20\n" )
   } else {
       /* 如果条件为 false 则执行以下语句 */
       fmt.Printf("a 不小于 20\n" )
   }
   fmt.Printf("a 的值为 : %d\n", a)
}
````

> D:\Program Files (x86)\Ampps\www\github\golang-study\4.if-for>go run ifelse.go

    a 不小于 20
    a 的值为 : 100

#### if 语句嵌套
    
if 或 else if 语句中嵌入一个或多个 if 或 else if 语句。
 
    if 布尔表达式 1 {
       /* 在布尔表达式 1 为 true 时执行 */
       if 布尔表达式 2 {
          /* 在布尔表达式 2 为 true 时执行 */
       }
    }

###### 实例

```go
package main

import "fmt"

func main() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200

	/* 判断条件 */
	if a == 100 {
		/* if 条件语句为 true 执行 */
		if b == 200 {
			/* if 条件语句为 true 执行 */
			fmt.Printf("a 的值为 100 ， b 的值为 200\n" )
		}
	}
	fmt.Printf("a 值为 : %d\n", a )
	fmt.Printf("b 值为 : %d\n", b )
}
```

> D:\Program Files (x86)\Ampps\www\github\golang-study\4.if-for>go run ifif.go

    a 的值为 100 ， b 的值为 200
    a 值为 : 100
    b 值为 : 200

#### switch 语句

switch 语句用于基于不同条件执行不同动作。

    switch var1 {
        case val1:
            ...
        case val2:
            ...
        default:
            ...
    }

###### 实例

```go
package main

import "fmt"

func main() {
	/* 定义局部变量 */
	var grade string = "B"
	var marks int = 90

	switch marks {
	case 90: grade = "A"
	case 80: grade = "B"
	case 50,60,70 : grade = "C"
	default: grade = "D"
	}

	switch {
	case grade == "A" :
		fmt.Printf("优秀!\n" )
	case grade == "B", grade == "C" :
		fmt.Printf("良好\n" )
	case grade == "D" :
		fmt.Printf("及格\n" )
	case grade == "F":
		fmt.Printf("不及格\n" )
	default:
		fmt.Printf("差\n" )
	}
	fmt.Printf("你的等级是 %s\n", grade )
}

```

> D:\Program Files (x86)\Ampps\www\github\golang-study\4.if-for>go run switch.go

    优秀!
    你的等级是 A

#### select 语句

select 语句类似于 switch 语句，但是select会随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。

    select {
        case communication clause  :
           statement(s);      
        case communication clause  :
           statement(s);
        /* 你可以定义任意数量的 case */
        default : /* 可选 */
           statement(s);
    }

###### 实例

```go
package main

import "fmt"

func main() {
	var c1, c2, c3 chan int
	var i1, i2 int

	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3):  // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}
}
```

> D:\Program Files (x86)\Ampps\www\github\golang-study\4.if-for>go run select.go

    no communication
