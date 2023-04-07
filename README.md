## the way to go

#### go

> Go是静态类型开发语言。具有自动垃圾回收（GC），丰富的内置类型（值类型、引用类型）, 函数多返回值，错误处理，匿名函数, 并发编程，反射，defer，简洁、安全、并行、开源等关键特征。语言层面支持并发，可充分利用CPU多核，Go语言编译程序可以媲美C或C++代码的速度。系统标准库功能完备，尤其是强大的网络库让建立Web服务成为再简单不过的事情。简单易学，内置runtime，支持继承、对象等，开发工具丰富，例如gofmt工具，自动格式化代码，让团队代码风格完美统一。同时也适合用来进行服务器编程，网络编程，包括Web应用、API应用，分布式编程等等。

> Go语言自2009年面世以来，已经有越来越多的公司开始转向Go语言开发，比如腾讯、百度、阿里、京东、小米以及360，而七牛云其技术栈基本上完全采用Go语言来开发。还有像今日头条、UBER这样的公司，他们也使用Go语言对自己的业务进行了彻底的重构。在全球范围内Go语言的使用不断增长，尤其是在云计算领域，用Go语言编写的几个主要云基础项目如**Docker**和**Kubernetes**，都取得了巨大成功。除此之外，还有各种有名的项目如etcd/consul/flannel等，均使用Go语言实现。

- 语法简洁。类C语言风格，简单易学，学习曲线平缓
- 代码风格统一。Go 语言提供了一套格式化工具——go fmt
- 开发效率高。开发效率与执行效率的完美结合，像写Python代码效率，编写C代码性能
- 天生支持并发
- 应用方向：区块链、分布式系统、云原生

#### [文档](http://docs.litao0501.top)

```
本地查看

git clone # 克隆项目

npm i docsify-cli -g # 如已安装 本步可忽略

docsify serve # 启动服务

http://localhost:3000
```

#### 目录

0. [环境搭建](0.install/README.md)
1. [hello world](1.demo/README.md)
2. [语言结构](2.structure/README.md)
3. [数据类型和声明](3.dataType/README.md)
4. [条件和循环](4.if-for/README.md)
5. [数组](5.arr/README.md)
6. [函数](6.fun/README.md)
7. [集合](7.map/README.md)
8. [范围](8.range/README.md)
9. [切片](9.slice/README.md)
10. [递归函数](10.recursive/README.md)
11. [接口](11.interface/README.md)
12. [错误处理](12.error/README.md)
13. [并发和互斥锁、读写锁](13.conc/README.md)
14. 包

	- [mysql](https://github.com/go-sql-driver/mysql)
	- [mysqlx ](github.com/jmoiron/sqlx)
	- [redis](https://github.com/go-redis/redis)
	- [gorm](https://github.com/go-gorm/gorm)

15. [依赖管理](15.mod/README.md)
16. [框架](16.frame/README.md)

	- [Iris](https://github.com/kataras/iris) 性能好，学习曲线低，教程丰富
	- [Beego](https://github.com/beego/beego) 简单易用，性能强大的企业级应用开发框架
	- [Buffalo](https://github.com/gobuffalo/buffalo) 使用 Go 语言快速构建 Web 应用
	- [Echo](https://github.com/labstack/echo) 简约高性能可扩展
	- [Gin](https://github.com/codegangsta/gin) Martini风格，但是性能要好很多
	- [GoFrame](https://github.com/gogf/gf) 是一款模块化、高性能、企业级的Go基础开发框架

17. 高级编程
