// Copyright 2023 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package render

import (
	"bytes"
	"os"
	"path/filepath"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

func (p *BookRendor) renderHomepage() error {
	markdown := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
	)

	// 已经有 html 文件
	if html, err := os.ReadFile(filepath.Join(p.Book.Root, "index.html")); err == nil {
		return os.WriteFile(
			filepath.Join(p.Book.Root, "book", "index.html"),
			html,
			0666,
		)
	}

	// 读取首页
	page_Content, err := os.ReadFile(filepath.Join(p.Book.Root, "index.md"))
	if err != nil {
		page_Content, err = os.ReadFile(filepath.Join(p.Book.Root, "readme.md"))
		if err != nil {
			page_Content, err = os.ReadFile(filepath.Join(p.Book.Root, "README.md"))
		}
		err = nil
	}

	var buf bytes.Buffer
	if err := markdown.Convert([]byte(page_Content), &buf); err != nil {
		return err
	}

	return os.WriteFile(
		filepath.Join(p.Book.Root, "book", "index.html"),
		buf.Bytes(),
		0666,
	)
}
