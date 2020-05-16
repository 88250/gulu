<p align = "center">
<img alt="Gulu" src="https://user-images.githubusercontent.com/873584/58315007-4100f080-7e43-11e9-9b10-b64a6a4a5d2d.png">
<br><br>
Go 语言常用工具库，这个轱辘还算圆！
<br><br>
<a title="Build Status" target="_blank" href="https://travis-ci.org/88250/gulu"><img src="https://img.shields.io/travis/88250/gulu.svg?style=flat-square"></a>
<a title="GoDoc" target="_blank" href="https://godoc.org/github.com/88250/gulu"><img src="http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square"></a>
<a title="Go Report Card" target="_blank" href="https://goreportcard.com/report/github.com/88250/gulu"><img src="https://goreportcard.com/badge/github.com/88250/gulu?style=flat-square"></a>
<a title="Coverage Status" target="_blank" href="https://coveralls.io/repos/github/88250/gulu/badge.svg?branch=master"><img src="https://img.shields.io/coveralls/github/88250/gulu.svg?style=flat-square&color=CC9933"></a>
<a title="Code Size" target="_blank" href="https://github.com/88250/gulu"><img src="https://img.shields.io/github/languages/code-size/88250/gulu.svg?style=flat-square"></a>
<br>
<a title="Apache License" target="_blank" href="https://github.com/88250/gulu/blob/master/LICENSE"><img src="https://img.shields.io/badge/license-apache2-orange.svg?style=flat-square"></a>
<a title="GitHub Commits" target="_blank" href="https://github.com/88250/gulu/commits/master"><img src="https://img.shields.io/github/commit-activity/m/88250/gulu.svg?style=flat-square"></a>
<a title="Last Commit" target="_blank" href="https://github.com/88250/gulu/commits/master"><img src="https://img.shields.io/github/last-commit/88250/gulu.svg?style=flat-square&color=FF9900"></a>
<a title="GitHub Pull Requests" target="_blank" href="https://github.com/88250/gulu/pulls"><img src="https://img.shields.io/github/issues-pr-closed/88250/gulu.svg?style=flat-square&color=FF9966"></a>
<a title="Hits" target="_blank" href="https://github.com/88250/hits"><img src="https://hits.b3log.org/88250/gulu.svg"></a>
<br><br>
<a title="GitHub Watchers" target="_blank" href="https://github.com/88250/gulu/watchers"><img src="https://img.shields.io/github/watchers/88250/gulu.svg?label=Watchers&style=social"></a>  
<a title="GitHub Stars" target="_blank" href="https://github.com/88250/gulu/stargazers"><img src="https://img.shields.io/github/stars/88250/gulu.svg?label=Stars&style=social"></a>  
<a title="GitHub Forks" target="_blank" href="https://github.com/88250/gulu/network/members"><img src="https://img.shields.io/github/forks/88250/gulu.svg?label=Forks&style=social"></a>  
<a title="Author GitHub Followers" target="_blank" href="https://github.com/88250"><img src="https://img.shields.io/github/followers/88250.svg?label=Followers&style=social"></a>
</p>

## 💡 简介

[Gulu](https://github.com/88250/gulu) 是一款 Go 语言常用工具库。

欢迎到 [Gulu 官方讨论区](https://hacpai.com/tag/gulu)了解更多。同时也欢迎关注 B3log 开源社区微信公众号 `B3log开源`：

![b3logos.png](https://img.hacpai.com/file/2019/10/image-d3c00d78.png)

## ✨ 功能

### 文件操作 `gulu.File`

* 获取文件大小
* 判断路径是否存在
* 判断文件是否是图片
* 按内容判断文件是否是可执行二进制
* 判断文件是否是目录
* 复制文件
* 复制目录

### Go 语言 `gulu.Go`

* 获取 Go API 源码目录路径
* 判断指定路径是否在 Go API 源码目录下
* 获取格式化工具名 ["gofmt", "goimports"]
* 获取 $GOBIN 下指定可执行程序名的绝对路径

### 日志记录 `gulu.Log`

* 提供可指定日志级别的日志记录器

### 网络相关 `gulu.Net`

* 获取本机第一张网卡的 IP 地址
* 获取本机第一张网卡的 MAC 地址

### 操作系统 `gulu.OS`

* 判断是否是 Windows
* 判断是否是 Linux
* 判断是否是 Darwin
* 获取当前进程的工作目录
* 获取用户 Home 目录路径

### panic 处理 `gulu.Panic`

* 包装 recover() 提供更好的报错日志格式

### 随机数 `gulu.Rand`

* 随机字符串
* 随机整数

### 返回值 `gulu.Ret`

* 提供普适返回值结构

### Rune `gulu.Rune`

* 判断 rune 是否为数字或字母
* 判断 rune 是否为字母

### 字符串 `gulu.Str`

* 字符串是否包含在字符串数组中
* 求最长公共子串

### Zip 压缩解压 `gulu.Zip`

* Zip 压缩和解压

## 🗃 案例

* [Pipe](https://github.com/88250/pipe)：一款小而美的博客平台，专为程序员设计
* [Wide](https://github.com/88250/wide)：一款基于 Web 的 Go 语言 IDE，随时随地玩 golang
* [BND](https://github.com/88250/baidu-netdisk-downloaderx)：一款图形界面的百度网盘不限速下载器，支持 Windows、Linux 和 Mac
* [协慌网](https://routinepanic.com)：专注编程问答汉化
* [链滴笔记](https://github.com/88250/liandi)：一款桌面端笔记应用，支持 Windows、Mac 和 Linux

## 💝 贡献

Gulu 肯定有一些不足之处：

* 代码不够优美
* 文档不够清晰
* 功能不够完善
* 可能存在缺陷
* ……

希望大家能和我们一起来完善该项目，无论是提交需求建议还是代码改进，我们都非常欢迎！

## 🏘️ 社区

* [讨论区](https://hacpai.com/tag/gulu)
* [报告问题](https://github.com/88250/gulu/issues/new/choose)

## 📄 授权

Gulu 使用 [木兰宽松许可证, 第2版](http://license.coscl.org.cn/MulanPSL2) 开源协议。

## 🙏 鸣谢

* [The Go Programming Language](https://golang.org)
