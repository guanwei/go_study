# go_study

开发环境为 MAC 系统

## Install Go

```
$ brew install go
```

检查 go 版本
```
$ go version
```

## VsCode 安装 go 插件

在 VsCode 中，使用快捷键：Command + Shift + P，然后键入 go:install/update tools，将所有 16 个插件都勾选上，然后点击 OK 即开始安装。

由于墙的原因，安装 go 插件会提示安装失败

### 解决方法

使用终端切换到 $GOPATH 下，按照下面目录结构来新建缺失的文件夹
```
src
├── github.com
|      └── golang
└── golang.org
       └── x
```

安装 tools 插件
```
$ cd $GOPATH/src/golang.org/x
$ git clone https://github.com/golang/tools.git tools
```

然后再安装其余插件
```
$ cd $GOPATH
$ go install github.com/ramya-rao-a/go-outline
$ go install github.com/acroca/go-symbols
$ go install golang.org/x/tools/cmd/guru
$ go install golang.org/x/tools/cmd/gorename
$ go install github.com/josharian/impl
$ go install github.com/rogpeppe/godef
$ go install github.com/sqs/goreturns
$ go install github.com/golang/lint/golint
$ go install github.com/cweill/gotests/gotests
$ go install github.com/ramya-rao-a/go-outline
$ go install github.com/acroca/go-symbols
$ go install golang.org/x/tools/cmd/guru
$ go install golang.org/x/tools/cmd/gorename
$ go install github.com/josharian/impl
$ go install github.com/rogpeppe/godef
$ go install github.com/sqs/goreturns
$ go install github.com/golang/lint/golint
$ go install github.com/cweill/gotests/gotests
```

### 设置 go 代码自动格式化

依次键入F1 -> Open User Settings 即可打开 vscode 用户配置文件，在文件末尾添加如下属性。
```
"go.formatTool": "goimports", //使用 goimports 工具进行代码格式化，或者使用 goreturns 和gofmt
```