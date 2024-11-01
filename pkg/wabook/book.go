// Copyright 2023 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wabook

import (
	"path/filepath"
)

type Book struct {
	Root    string
	Info    *BookToml
	Summary *Summary
	Talks   []string
}

func LoadBook(path string) (book *Book, err error) {
	bookIniPath := filepath.Join(path, "book.ini")
	bookTomlPath := filepath.Join(path, "book.toml")

	book = &Book{Root: path}
	book.Info, err = LoadConfig(bookIniPath)
	if err != nil {
		if info, errx := LoadConfig(bookTomlPath); errx != nil {
			if fileExists(bookIniPath) || fileExists(bookTomlPath) {
				// 文件存在说明有错误
				return nil, err
			} else {
				// 不存在则生成一个空的
				book.Info = &BookToml{}
			}
		} else {
			book.Info = info
		}
	}
	book.Summary, err = LoadSummary(filepath.Join(path, "SUMMARY.md"))
	if err != nil {
		if fileExists(filepath.Join(path, "SUMMARY.md")) {
			return nil, err
		} else {
			err = nil
		}
		book.Summary = &Summary{}
	}
	book.Talks = loadTalks(book.Root)
	return
}
