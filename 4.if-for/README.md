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

----

## 循环语句

#### for 循环

重复执行语句块。

Go 语言的 For 循环有 3 种形式，只有其中的一种使用分号。

和 C 语言的 for 一样：

    for init; condition; post { }
    
和 C 的 while 一样：

    for condition { }
    
和 C 的 for(;;) 一样：

    for { }

###### 实例

计算 1 到 100 的数字之和：

```go
package main

import "fmt"

func main() {
        sum := 0
        for i := 0; i <= 100; i++ {
                sum += i
        }
        fmt.Println(sum)
}
```

> D:\Program Files (x86)\Ampps\www\github\golang-study\4.if-for>go run for.go

    5050

#### 循环嵌套

在 for 循环中嵌套一个或多个 for 循环。

    for [condition |  ( init; condition; increment ) | Range]
    {
       for [condition |  ( init; condition; increment ) | Range]
       {
          statement(s);
       }
       statement(s);
    }
    
###### 实例

循环嵌套来输出 2 到 100 间的素数：

```go
package main

import "fmt"

func main() {
	/* 定义局部变量 */
	var i, j int

	for i=2; i < 100; i++ {
		for j=2; j <= (i/j); j++ {
			if i%j == 0 {
				break // 如果发现因子，则不是素数
			}
		}
		if j > (i/j) {
			fmt.Printf("%d  是素数\n", i)
		}
	}
}

```

> D:\Program Files (x86)\Ampps\www\github\golang-study\4.if-for>go run forFor.go

    2  是素数
    3  是素数
    5  是素数
    7  是素数
    11  是素数
    13  是素数
    17  是素数
    19  是素数
    23  是素数
    29  是素数
    31  是素数
    37  是素数
    41  是素数
    43  是素数
    47  是素数
    53  是素数
    59  是素数
    61  是素数
    67  是素数
    71  是素数
    73  是素数
    79  是素数
    83  是素数
    89  是素数
    97  是素数

#### 循环控制

- break 语句：经常用于中断当前 for 循环或跳出 switch 语句。

- continue 语句： 跳过当前循环的剩余语句，然后继续进行下一轮循环。

- goto 语句：将控制转移到被标记的语句。

#### 无线循环

如果循环中条件语句永远不为 false 则会进行无限循环，我们可以通过 for 循环语句中只设置一个条件表达式来执行无限循环：

```go
package main

import "fmt"

func main() {
    for true  {
        fmt.Printf("这是无限循环。\n");
    }
}
```

[首页](../README.md)

[下一章](../5.arr/README.md)
