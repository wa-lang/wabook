<div align="center">
<h1>waBook: Create book/talk from markdown files.</h1>

[简体中文](https://github.com/wa-lang/wabook/blob/master/README-zh.md) | [English](https://github.com/wa-lang/wabook/blob/master/README.md) 


</div>
<div align="center">

[![Build Status](https://github.com/wa-lang/wabook/actions/workflows/wabook.yml/badge.svg)](https://github.com/wa-lang/wabook/actions/workflows/wabook.yml)
[![GitHub release](https://img.shields.io/github/v/tag/wa-lang/wabook.svg?label=release)](https://github.com/wa-lang/wabook/releases)
[![license](https://img.shields.io/github/license/wa-lang/wa.svg)](https://github.com/wa-lang/wa/blob/master/LICENSE)

</div>

## Features

- Create book from Markdown
- Create slide from Markdown
- Create book discuss based on Github Discuss
- Supports custom header/footer

## Example

Book list built using [waBook](https://github.com/wa-lang/wabook):

- 《Go语言圣经》: https://gopl-zh.github.io
- 《Go语言高级编程》: https://github.com/chai2010/advanced-go-programming-book
- 《Go语言定制指南》: https://github.com/chai2010/go-ast-book
- 《µGo语言实现(从头开发一个迷你Go语言编译器)》: https://github.com/wa-lang/ugo-compiler-book
- 《凹语言手册》: https://wa-lang.org/man/
- 《Wa-lang's Manual》: https://wa-lang.github.io/man/en/
- 《VS Code插件开发》: https://chai2010.cn/vscode-extdev-book/
- 《Go语言圣经读书笔记》: https://github.com/chai2010/gopl-notes-zh

Slide list built using [waBook](https://github.com/wa-lang/wabook):

- 凹语言map与Φ指令的纠葛: https://wa-lang.org/talks/ssa-bug/
- Go编译器定制简介: https://wa-lang.github.io/ugo-compiler-book/go-compiler-intro.html

## Install

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

- init: Initialize a basic version of Book
- build: Build the Markdown Book into html
- serve: Build and start the service to facilitate local viewing of the effect
- clean: Delete the built book subdirectory

## `book.ini` file

No annotations supported, no undefined properties supported:

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

## `SUMMARY.md` file

```md
# Summary

[Preface](preface.md)

- [Chapter 1](./src/chapter_1.md)
  - [Chapter 1.1](./src/chapter_1.1.md)
  - [Chapter 1.2](./src/chapter_1.2.md)

- [Chapter 2](./src/chapter_2.md)

<!-- comment -->
```

## Markdown file

```md
# Chapter 1

[Github Repo](https://github.com/wa-lang/wabook): `[Github Repo](https://github.com/wa-lang/wabook)`


Image: `![](../images/video-001.png)`:

![](../images/video-001.png)

OK!
```

Do not support inline HTML。
