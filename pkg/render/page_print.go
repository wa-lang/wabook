// Copyright 2023 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package render

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/text"
)

func (p *BookRendor) renderPrintPage() error {
	var b strings.Builder
	for i, page := range p.PageInfos {
		s, err := p.renderMarkdownString(i, page)
		if err != nil {
			return err
		}
		b.WriteString(s)
	}

	// todo
	printInfo := &StaticPageInfo{
		Content: b.String(),
	}

	// render to page
	var fnMap = template.FuncMap{
		"genSidebarItems": func(pageRoot string, pageIndex int) string {
			return SidebarTree(p.SidebarItems).GenHTML(pageRoot, pageIndex)
		},
	}

	var buf bytes.Buffer
	t := template.Must(template.New("").Funcs(fnMap).Parse(tmplPage))
	err := t.Execute(&buf, printInfo)
	if err != nil {
		return err
	}

	dstAbsPath := filepath.Join(p.Book.Root, "book/print.html")

	os.MkdirAll(filepath.Dir(dstAbsPath), 0777)
	if err := os.WriteFile(dstAbsPath, buf.Bytes(), 0666); err != nil {
		return err
	}

	return nil
}

func (p *BookRendor) renderMarkdownString(idx int, page *PageInfo) (string, error) {
	markdown := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
	)

	reader := text.NewReader([]byte(page.Content))
	doc := markdown.Parser().Parse(reader)
	p.fixupLinkPath(doc, page)

	var b strings.Builder
	if err := markdown.Renderer().Render(&b, []byte(page.Content), doc); err != nil {
		return "", err
	}

	return b.String(), nil
}

func (p *BookRendor) fixupLinkPath(n ast.Node, page *PageInfo) {
	ast.Walk(n, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if entering {
			switch n := n.(type) {
			case *ast.Image:
				var dest = string(n.Destination)
				for strings.HasPrefix(dest, "../") {
					dest = strings.TrimPrefix(dest, "../")
				}
				n.Destination = []byte(dest)
			case *ast.Link:
				var dest = string(n.Destination)
				for strings.HasPrefix(dest, "../") {
					dest = strings.TrimPrefix(dest, "../")
				}
				if ext := filepath.Ext(dest); strings.EqualFold(ext, ".md") {
					dest = dest[:len(dest)-len(".md")] + ".html"
				}
				n.Destination = []byte(dest)
			}
		}
		return ast.WalkContinue, nil
	})
}
