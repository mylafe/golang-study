## 依赖管理

> **go module**依赖管理工具

#### 环境变量

    go env
    GO111MODULE

    GO111MODULE=off禁用模块支持，编译时会从GOPATH和vendor文件夹中查找包。
    GO111MODULE=on启用模块支持，编译时会忽略GOPATH和vendor文件夹，只根据 go.mod下载依赖。
    GO111MODULE=auto，当项目在$GOPATH/src外且项目根目录有go.mod文件时，开启模块支持。


#### GOPROXY

    go env -w GOPROXY=https://goproxy.cn,direct

#### go mod

```
go mod download    下载依赖的module到本地cache（默认为$GOPATH/pkg/mod目录）
go mod edit        编辑go.mod文件
go mod graph       打印模块依赖图
go mod init        初始化当前文件夹, 创建go.mod文件
go mod tidy        增加缺少的module，删除无用的module
go mod vendor      将依赖复制到vendor下
go mod verify      校验依赖
go mod why         解释为什么需要依赖

go get -u将会升级到最新的次要版本或者修订版本(x.y.z, z是修订版本号， y是次要版本号)
go get -u=patch将会升级到最新的修订版本
go get package[@version](https://github.com/version "@version")将会升级到指定的版本号version 
```

[首页](../README.md)

[下一章](../16.frame/README.md)
