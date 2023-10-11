// Copyright 2023 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mnbook

import (
	"path/filepath"
)

type Book struct {
	Root    string
	Info    *BookToml
	Summary *Summary
}

func LoadBook(path string) (book *Book, err error) {
	book = &Book{Root: path}
	book.Info, err = LoadConfig(filepath.Join(path, "book.toml"))
	if err != nil {
		return nil, err
	}
	book.Summary, err = LoadSummary(filepath.Join(path, "SUMMARY.md"))
	if err != nil {
		return nil, err
	}
	return
}
