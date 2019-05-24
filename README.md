# gulu

## 简介

Go 语言常用工具库，好轱辘大家用。

## 功能

### `gulu.File`

* 获取文件大小
* 判断路径是否存在
* 判断文件是否是图片
* 按内容判断文件是否是可执行二进制
* 判断文件是否是目录
* 复制文件
* 复制目录

### `gulu.Go`

* 获取 Go API 源码目录路径
* 判断指定路径是否在 Go API 源码目录下
* 获取格式化工具名 ["gofmt", "goimports"]
* 获取 $GOBIN 下指定可执行程序名的绝对路径

### `gulu.NewLogger()`

* 提供可指定日志级别的日志记录器

### `gulu.Net`

* 获取本机第一张网卡的地址

### `gulu.OS`

* 判断是否是 Windows
* 获取当前进程的工作目录
* 获取用户 Home 目录路径

### `gulu.Recover()`

* 包装 recover() 提供更好的报错日志格式

### `gulu.Rand`

* 随机字符串
* 随机整数

### `gulu.NewResult()`

* 提供普适返回值结构

### `gulu.String`

* 字符串是否包含在字符串数组中
* 求最长公共子串

### `gulu.Zip`

* Zip 压缩和解压

## 案例

* [Pipe](https://github.com/b3log/pipe)：一款小而美的博客平台，专为程序员设计
* [Wide](https://github.com/b3log/wide)：一款基于 Web 的 Go 语言 IDE，随时随地玩 golang
* [协慌网](https://routinepanic.com)：专注编程问答汉化

如果你也在使用 Gulu，欢迎通过 PR 将你的项目添加到这里。

## 贡献

Gulu 肯定有一些不足之处：

* 实现存在缺陷
* 代码不够优美
* 文档不够清晰
* 功能不够完善
* ……

欢迎大家和我们一起来完善该项目，无论是提交需求建议还是代码改进，我们都非常欢迎！

## 社区

* [讨论区](https://hacpai.com/tag/gulu)
* [报告问题](https://github.com/b3log/gulu/issues/new/choose)

## 授权

Gulu 使用 [Apache License, Version 2](https://www.apache.org/licenses/LICENSE-2.0) 开源协议。

## 鸣谢

* [The Go Programming Language](https://golang.org)
