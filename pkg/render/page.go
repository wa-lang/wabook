// Copyright 2023 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package render

import (
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"

	"github.com/wa-lang/mnbook/pkg/mnbook"
)

//go:embed tmpl/page.html
var tmplPage string

//go:embed tmpl/present.html
var tmplPresent string

//go:embed tmpl/action.tmpl
var tmplAction string

//go:embed tmpl/print.html
var tmplPrintPage string

func (p *BookRendor) run(book *mnbook.Book) (err error) {
	if err := p.init(book); err != nil {
		return err
	}

	if err := p.renderStaticFile(); err != nil {
		return err
	}
	if err := p.renderStaticAsset(); err != nil {
		return err
	}

	if err := p.renderHomepage(); err != nil {
		return err
	}

	if err := p.renderAllPages(); err != nil {
		return err
	}

	if err := p.renderAllTalkPages(); err != nil {
		return err
	}

	return nil
}

func (p *BookRendor) init(book *mnbook.Book) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()

	*p = BookRendor{}
	p.Book = book
	if len(p.Ignores) == 0 {
		p.Ignores = append(p.Ignores, DefaultBookIgnores...)
	}

	p.BookInfo = &BookInfo{
		Title:           book.Info.Book.Title,
		GitRepoIcon:     book.Info.OutputHtml.GitRepositoryIcon,
		GitRepoUrl:      book.Info.OutputHtml.GitRepositoryUrl,
		EditUrlTemplate: book.Info.OutputHtml.EditUrlTemplate,
		Giscus:          book.Info.Giscus,
	}

	sidebarNumbers := p.buildSidebarNumbers()
	for i, item := range p.Book.Summary.Chapters {
		sidebarItem := p.buildSidebarItem(i, item, sidebarNumbers)
		p.SidebarItems = append(p.SidebarItems, sidebarItem)
	}

	for i, item := range p.Book.Summary.Chapters {
		pageInfo := p.buildPageInfo(i, item)
		pageInfo.SidebarItems = p.SidebarItems
		pageInfo.EditUrlTemplate = strings.ReplaceAll(
			p.BookInfo.EditUrlTemplate, "{path}", pageInfo.Path,
		)

		p.PageInfos = append(p.PageInfos, pageInfo)
	}

	return
}

func (p *BookRendor) renderAllPages() error {
	for i, page := range p.PageInfos {
		if err := p.renderPage(i, page); err != nil {
			return err
		}
	}
	if err := p.renderPrintPage(); err != nil {
		return nil
	}
	if err := p.render404Page(); err != nil {
		return nil
	}
	return nil
}

func (p *BookRendor) renderPage(idx int, page *PageInfo) error {
	markdown := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
	)

	var buf bytes.Buffer
	if err := markdown.Convert([]byte(page.Content), &buf); err != nil {
		return err
	}

	// md -> raw html
	page.RawHtml = buf.String()
	buf.Reset()

	// render to page
	var fnMap = template.FuncMap{
		"genSidebarItems": func(pageRoot string, pageIndex int) string {
			return SidebarTree(p.SidebarItems).GenHTML(pageRoot, pageIndex)
		},
	}

	t := template.Must(template.New("").Funcs(fnMap).Parse(tmplPage))
	err := t.Execute(&buf, page)
	if err != nil {
		return err
	}

	dstAbsPath := filepath.Join(p.Book.Root, "book", page.Path)
	if ext := filepath.Ext(dstAbsPath); strings.EqualFold(ext, ".md") {
		dstAbsPath = dstAbsPath[:len(dstAbsPath)-len(".md")]
	}
	dstAbsPath += ".html"

	os.MkdirAll(filepath.Dir(dstAbsPath), 0777)
	if err := os.WriteFile(dstAbsPath, buf.Bytes(), 0666); err != nil {
		return err
	}

	return nil
}

func (p *BookRendor) buildSidebarNumbers() []string {
	var numbers = make([]string, len(p.Book.Summary.Chapters))
	var idxCountor [4]int
	var prevIdx int

	for i, item := range p.Book.Summary.Chapters {
		if strings.TrimSpace(item.Prefix) == "" {
			numbers[i] = "-skip-"
			continue
		}

		k := len(item.Prefix) / 2
		switch {
		case k == prevIdx:
			if k < len(idxCountor) {
				idxCountor[k]++
			}
		case k > prevIdx:
			for j := prevIdx + 1; j < len(idxCountor); j++ {
				idxCountor[j] = 0
			}
			if k < len(idxCountor) {
				idxCountor[k]++
			}
		case k < prevIdx:
			for j := k + 1; j < len(idxCountor); j++ {
				idxCountor[j] = 0
			}
			if k < len(idxCountor) {
				idxCountor[k]++
			}
		}

		prevIdx = k

		var endIdx = len(idxCountor) - 1
		for endIdx >= 0 && idxCountor[endIdx] == 0 {
			endIdx--
		}

		var sb strings.Builder
		for k := 0; k <= endIdx; k++ {
			sb.WriteString(strconv.Itoa(idxCountor[k]))
			sb.WriteRune('.')
		}
		numbers[i] = sb.String()
	}

	for i := 0; i < len(numbers); i++ {
		if numbers[i] == "-skip-" {
			numbers[i] = ""
			continue
		}
	}

	return numbers
}

func (p *BookRendor) buildSidebarItem(idx int, item mnbook.SummaryItem, sidebarNumbers []string) *SidebarItem {
	sidebarItem := new(SidebarItem)
	sidebarItem.Prefix = item.Prefix
	sidebarItem.Number = sidebarNumbers[idx]
	sidebarItem.Name = item.Name
	sidebarItem.Location = p.getPageRelPath(item)

	if ext := filepath.Ext(sidebarItem.Location); strings.EqualFold(ext, ".md") {
		sidebarItem.Location = sidebarItem.Location[:len(sidebarItem.Location)-len(".md")] + ".html"
	}
	return sidebarItem
}

func (p *BookRendor) buildPageInfo(idx int, item mnbook.SummaryItem) *PageInfo {
	page := new(PageInfo)

	page.Index = idx
	page.Book = p.BookInfo

	page.Root = p.getPageRootPath(item)
	page.Path = p.getPageRelPath(item)
	page.Title = p.getPageTitle(item)
	page.Content = p.loadPageContent(item)

	if idx > 0 {
		page.PrevUrl = p.SidebarItems[idx-1].Location
		page.HasPrev = true
	}

	if idx < len(p.Book.Summary.Chapters)-1 {
		page.NextUrl = p.SidebarItems[idx+1].Location
		page.HasNext = true
	}

	return page
}

func (p *BookRendor) getPageLocalAbsPath(item mnbook.SummaryItem) string {
	return filepath.Clean(filepath.Join(p.Book.Root, p.Book.Info.Book.Src, item.Location))
}

func (p *BookRendor) getPageRelPath(item mnbook.SummaryItem) string {
	absPath := filepath.Clean(filepath.Join(p.Book.Root, p.Book.Info.Book.Src, item.Location))
	relPath, err := filepath.Rel(p.Book.Root, absPath)
	if err != nil {
		panic(fmt.Sprintf("BookRendor.getItemRelPath(%+v): %v", item, err))
	}
	return relPath
}

func (p *BookRendor) getPageRootPath(item mnbook.SummaryItem) string {
	absPath := filepath.Clean(filepath.Join(p.Book.Root, p.Book.Info.Book.Src, filepath.Dir(item.Location)))
	relPath, err := filepath.Rel(absPath, p.Book.Root)

	if err != nil {
		panic(fmt.Sprintf("BookRendor.getItemRoot(%+v): %v", item, err))
	}
	return relPath
}

func (p *BookRendor) getPageTitle(item mnbook.SummaryItem) string {
	return fmt.Sprintf("%s - %s", item.Name, p.BookInfo.Title)
}

func (p *BookRendor) loadPageContent(item mnbook.SummaryItem) string {
	localAbsPath := p.getPageLocalAbsPath(item)
	if _, err := os.Lstat(localAbsPath); errors.Is(err, os.ErrNotExist) {
		var buf bytes.Buffer
		fmt.Fprintln(&buf, "#", item.Name)
		fmt.Fprintln(&buf, "")
		fmt.Fprintln(&buf, "TODO")
		fmt.Fprintln(&buf, "")

		os.MkdirAll(filepath.Dir(localAbsPath), 0777)
		os.WriteFile(localAbsPath, buf.Bytes(), 0666)
	}

	data, err := os.ReadFile(localAbsPath)
	if err != nil {
		panic(fmt.Sprintf("BookRendor.loadPageContent(%+v): %v", item, err))
	}
	return string(data)
}
