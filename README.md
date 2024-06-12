# Mini Markdown Book

简单的 Markdown 图书构建工具。

使用 [MNBook](https://github.com/wa-lang/mnbook) 构建的图书列表:

- 《凹语言手册》: https://wa-lang.org/man/
- 《Wa-lang's Manual》: https://wa-lang.github.io/man/en/
- 《VS Code插件开发》: https://chai2010.cn/vscode-extdev-book/
- 《Go语言圣经读书笔记》: https://github.com/chai2010/gopl-notes-zh

## 安装

```
$ go install github.com/wa-lang/mnbook@latest
```

## 命令行

```
$ mnbook
NAME:
   mnbook - A tool for build mini markdown book

USAGE:
   mnbook [global options] command [command options] [arguments...]

COMMANDS:
   init   Creates a new book
   build  Builds a book from its markdown files
   serve  Serves a book at http://localhost:3000
   clean  Deletes a built book

GLOBAL OPTIONS:
   --help, -h  show help

 See "https://github.com/wa-lang/mnbook" for more information.
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
git-repository-url = "https://github.com/wa-lang/mnbook"
edit-url-template = "https://github.com/wa-lang/mnbook/edit/master/testdata/{path}"
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

[Github Repo](https://github.com/wa-lang/mnbook): `[Github Repo](https://github.com/wa-lang/mnbook)`


Image: `![](../images/video-001.png)`:

![](../images/video-001.png)

OK!
```

不支持内联 HTML。
