<div align="center">
<h1>waBook: 简单的 Markdown 图书构建工具.</h1>

[简体中文](https://github.com/wa-lang/wabook/blob/master/README-zh.md) | [English](https://github.com/wa-lang/wabook/blob/master/README.md) 


</div>
<div align="center">

[![Build Status](https://github.com/wa-lang/wabook/actions/workflows/wabook.yml/badge.svg)](https://github.com/wa-lang/wabook/actions/workflows/wabook.yml)
[![GitHub release](https://img.shields.io/github/v/tag/wa-lang/wabook.svg?label=release)](https://github.com/wa-lang/wabook/releases)
[![license](https://img.shields.io/github/license/wa-lang/wa.svg)](https://github.com/wa-lang/wa/blob/master/LICENSE)

</div>

## 特性

- 支持 Markdown 格式的电子书构建
- 支持 Markdown 格式的幻灯片构建
- 支持基于 Github Discuss 的留言功能
- 电子书支持自定义 页眉/页脚

## 案例

使用 [waBook](https://github.com/wa-lang/wabook) 构建的图书列表:

- 《Go语言圣经》: https://gopl-zh.github.io
- 《Go语言高级编程》: https://github.com/chai2010/advanced-go-programming-book
- 《Go语言定制指南》: https://github.com/chai2010/go-ast-book
- 《µGo语言实现(从头开发一个迷你Go语言编译器)》: https://github.com/wa-lang/ugo-compiler-book
- 《凹语言手册》: https://wa-lang.org/man/
- 《Wa-lang's Manual》: https://wa-lang.github.io/man/en/
- 《VS Code插件开发》: https://chai2010.cn/vscode-extdev-book/
- 《Go语言圣经读书笔记》: https://github.com/chai2010/gopl-notes-zh

使用 [waBook](https://github.com/wa-lang/wabook) 构建的幻灯片列表:

- 凹语言map与Φ指令的纠葛: https://wa-lang.org/talks/ssa-bug/
- Go编译器定制简介: https://wa-lang.github.io/ugo-compiler-book/go-compiler-intro.html


## 安装

```
$ go install github.com/wa-lang/wabook@latest
```

## 命令行

```
$ wabook
NAME:
   wabook - A tool for build mini markdown book

USAGE:
   wabook [global options] command [command options] [arguments...]

COMMANDS:
   init   Creates a new book
   build  Builds a book from its markdown files
   serve  Serves a book at http://localhost:3000
   clean  Deletes a built book

GLOBAL OPTIONS:
   --help, -h  show help

 See "https://github.com/wa-lang/wabook" for more information.
```

- init: 初始化一个 Book 基础版本
- build: 将 Markdown 的 Book 构建为 html
- serve: 构建并启动服务, 方便本地查看效果
- clean: 删除构建的 book 子目录

## `book.ini` 文件

不支持注释，不支持未定义属性：

```ini
[book]
authors = ["chai2010"]
description = ""
language = "zh"
src = "."
title = "book title"

[output.html]
git-repository-icon = "fa-github"
git-repository-url = "https://github.com/wa-lang/wabook"
edit-url-template = "https://github.com/wa-lang/wabook/edit/master/testdata/{path}"
```

## `SUMMARY.md` 文件

```md
# Summary

[Preface](preface.md)

- [Chapter 1](./src/chapter_1.md)
  - [Chapter 1.1](./src/chapter_1.1.md)
  - [Chapter 1.2](./src/chapter_1.2.md)

- [Chapter 2](./src/chapter_2.md)

<!-- comment -->
```

## Markdown 文件

```md
# Chapter 1

[Github Repo](https://github.com/wa-lang/wabook): `[Github Repo](https://github.com/wa-lang/wabook)`


Image: `![](../images/video-001.png)`:

![](../images/video-001.png)

OK!
```

不支持内联 HTML。
