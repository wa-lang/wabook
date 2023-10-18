// Copyright 2023 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package render

import "github.com/wa-lang/mnbook/pkg/mnbook"

var DefaultBookIgnores = []string{
	".git", "_git",
	"book", "book.toml",
	"node_modules",
}

type BookRendor struct {
	Book         *mnbook.Book
	BookInfo     *BookInfo
	PageInfos    []*PageInfo
	SidebarItems []*SidebarItem
	Ignores      []string
}

type BookInfo struct {
	Title           string
	GitRepoIcon     string
	GitRepoUrl      string
	EditUrlTemplate string
}

type PageInfo struct {
	Index int

	Book            *BookInfo
	SidebarItems    []*SidebarItem
	EditUrlTemplate string

	Root  string
	Path  string
	Title string

	Content string
	RawHtml string

	HasPrev bool
	HasNext bool

	PrevUrl string
	NextUrl string
}

type SidebarItem struct {
	Prefix   string
	Number   string // 1.1, 1.2, ...
	Name     string
	Location string
}

func New() *BookRendor {
	return &BookRendor{
		Ignores: append([]string{}, DefaultBookIgnores...),
	}
}

func (p *BookRendor) Run(book *mnbook.Book) error {
	return p.run(book)
}
