# 第一章 入门

运行单个go文件
```
$ go run helloworld.go
Hello, 世界
```

也可以是网络路径
```
$ go run gopl.io/ch1/helloworld
Hello,世界
```

编译单个go文件，会在当前所在目录下生成可执行的二进制文件
```
$ go build helloworld.go
$ ./helloworld
Hello, 世界
```

也可以是网络路径
```
$ go build gopl.io/ch1/helloworld
$ ./helloworld
Hello, 世界
```

许多文本编辑器可以配置为每次在保存文件时自动运行gofmt，因此源文件总是可以保持正确的形式。
此外，一个相关的工具goimports可以按需管理导入声明的插入和移除。

执行下面的命令安装 goimports
```
$ go get golang.org/x/tools/cmd/goimports
```

## 命令行参数 os.Args

os.Args 的第一个元素是 os.Args[0]，它是命令本身的名字；另外的元素是程序开始执行时的参数 os.Args[1:]

注释以//开头，习惯上，在一个包声明前，使用注释对其进行描述；对于main包，注释是一个或多个完整的句子，用来对这个程序进行整体概括。

变量可以在声明的时候初始化。如果变量没有明确地初始化，它将隐式地初始化为这个类型的空值。数字初始化为0，字符串初始化为""。

for 是 Go 里面的唯一循环语句。它有几种形式：
```
for initialization; condition; post {
    // 零个或多个语句
}

// 传统的 "while" 循环
for condition {
    // ...
}

// 传统的无限循环
for {
    // ...
}
```
无限循环可以通过break或return等语句终止。

```
for _, arg := range os.Args[1:] {
    // ...
}
```
每一次迭代，range产生一对值：索引和这个索引处元素的值。
Go不允许存在无用的临时变量，不然会出现编译错误。解决方案是使用`空标识符`，它的名字是_(即下划线)。

以下几种声明字符串变量的方式是等价的：
```
s := ""
var s string
var s = ""
var s string = ""
```

通过 += 语句追加旧的字符串，会生成一个新的字符串，旧的内容就不再使用，会被例行垃圾回收。如果有大量的数据需要处理，这样代价会比较大。一个简单和高效的方式是使用 strings 包中的 Join 函数。

## 找出重复行

map 存储一个键/值对集合，内置函数 make 可以用来新建 map，它还可以有其他用途。
语句 counts[input.Text()]++ 等价于下面的两个语句：
```
line := input.Text()
counts[line] = counts[line] + 1
```

函数 fmt.Printf 从一个表达式列表生成格式化的输出。
```
verb            描述
%d              十进制整数
%x, %o, %b      十六进制、八进制、二进制整数
%f, %g, %e      浮点数：如 3.141593， 3.141592653589793， 3.141593e+00
%t              布尔型：true 或 false
%c              字符（Unicode码点）
%s              字符串
%q              带引号字符串（如"abc"）或者字符（'c'）
%v              内置格式的任何值
%T              任何值的类型
%%              百分号本身（无操作数）
```

函数 os.Open 返回两个值。第一个是打开的文件 (*os.File)，该文件随后被Scanner读取。第二个返回值是一个内置的error类型的值。如果err不是nil，说明出错了。这时，error的值描述错误的原因。

ioutil.ReadFile 函数读取整个命名文件的内容，strings.Split 函数将一个字符串分割为一个由子串组成的slice。

const声明用来给常量命名，常量是其值在编译期间固定的量。常量必须是数字、字符串或布尔值。

结构体由一组称为字段的值组成，字段通常有不同的数据类型，他们一起组成单个对象，作为一个单元被对待。

`rand.Seed(time.Now().UnixNano())` 设置随机数种子，加上这行代码，可以保证每次随机都是随机的。

goroutine是一个并发执行函数。通道是一种允许某一例程向另一个例程传递指定类型的值的通信机制。

确保最多只有一个goroutine在同一时间访问变量，这正是`mu.Lock()`和`mu.Unlock()`语句的作用。

使用strconv.Atoi函数可以将字符串参数转换成整型。可以通过`go doc strconv.Atoi`来查看文档。