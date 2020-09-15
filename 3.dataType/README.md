#### 数据类型和声明

#### 数据类型

###### 字符串 string

只能用一对双引号（""）或反引号（``）括起来定义，不能用单引号（''）定义！

###### 布尔 bool

只有 true 和 false，默认为 false。

###### 数字 整型

int8 uint8

int16 uint16

int32 uint32

int64 uint64

int uint，具体长度取决于 CPU 位数。

###### 浮点型

float32 float64

#### 常量声明

常量，在程序编译阶段就确定下来的值，而程序在运行时无法改变该值。

###### 单个常量声明

第一种：const 变量名称 数据类型 = 变量值

如果不赋值，使用的是该数据类型的默认值。

第二种：const 变量名称 = 变量值

根据变量值，自行判断数据类型。

###### 多个常量声明

第一种：const 变量名称,变量名称 ... ,数据类型 = 变量值,变量值 ...

第二种：const 变量名称,变量名称 ... = 变量值,变量值 ...

###### 实例

    constant.go
    
> D:\Program Files (x86)\Ampps\www\github\golang-study\3.dataType>go run constant.go
  
    mylafe
    30
    mylafe wave
    mylafe 30

#### 变量声明

###### 单个变量声明

第一种：var 变量名称 数据类型 = 变量值

如果不赋值，使用的是该数据类型的默认值。

第二种：var 变量名称 = 变量值

根据变量值，自行判断数据类型。

第三种：变量名称 := 变量值

省略了 var 和数据类型，变量名称一定要是未声明过的。

###### 多个变量声明

第一种：var 变量名称,变量名称 ... ,数据类型 = 变量值,变量值 ...

第二种：var 变量名称,变量名称 ... = 变量值,变量值 ...

第三种：变量名称,变量名称 ... := 变量值,变量值 ...

###### 实例

    variable.go
    
> D:\Program Files (x86)\Ampps\www\github\golang-study\3.dataType>go run variable.go

    31 32 33
    31 32 33
    mylafe 30
    wave true 180.66

#### 打印

fmt.Print：输出到控制台（仅只是输出）

fmt.Println：输出到控制台并换行

fmt.Printf：仅输出格式化的字符串和字符串变量（整型和整型变量不可以）

fmt.Sprintf：格式化并返回一个字符串，不输出。

###### 实例

    print.go
    
> D:\Program Files (x86)\Ampps\www\github\golang-study\3.dataType>go run print.go

    输出到控制台不换行---
    输出到控制台并换行
    name=Tom,age=30
    name=Tom,age=30,height=180.57
 