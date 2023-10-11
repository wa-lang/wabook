// Copyright 2023 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package render

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

type SidebarTree []*SidebarItem

func (p SidebarTree) GenHTML(pageRoot string, pageIndex int) string {
	var buf bytes.Buffer
	p.genHTML(&buf, pageRoot, pageIndex, 0)
	return buf.String()
}

func (p SidebarTree) genHTML(w io.Writer, pageRoot string, pageIndex int, idx int) {
	if idx >= len(p) {
		return
	}

	var indent = p.getItemIndent(idx)
	var prefix = strings.Repeat(" ", indent)

	if indent == 0 {
		fmt.Fprintf(w, "%s%s\n", prefix, `<ol class="chapter">`)
		defer fmt.Fprintf(w, "%s%s\n", prefix, `</ol>`)
	} else {
		fmt.Fprintf(w, "%s%s\n", prefix, `<ol class="section">`)
		defer fmt.Fprintf(w, "%s%s\n", prefix, `</ol>`)
	}

	for _, itemIdx := range p.Siblings(idx) {
		var liContent string
		var calssActive string
		if itemIdx == pageIndex {
			calssActive = `class="active"`
		}

		if item := p[itemIdx]; item.Prefix == "" {
			liContent = fmt.Sprintf(
				`<a href="%s/%s" %s>%s</a>`,
				pageRoot,
				item.Location,
				calssActive,
				item.Name,
			)
		} else {
			liContent = fmt.Sprintf(
				`<a href="%s/%s" %s><strong aria-hidden="true">%s</strong> %s</a>`,
				pageRoot,
				item.Location,
				calssActive,
				item.Number,
				item.Name,
			)
		}

		fmt.Fprintf(w, "%s  %s\n", prefix, `<li class="chapter-item expanded ">`)
		fmt.Fprintf(w, "%s    %s\n", prefix, liContent)
		fmt.Fprintf(w, "%s  %s\n", prefix, `</li>`)

		if len(p.Children(itemIdx)) > 0 {
			p.genHTML(w, pageRoot, pageIndex, itemIdx+1)
		}
	}
}

func (p SidebarTree) Siblings(idx int) []int {
	if idx >= len(p) {
		return nil
	}

	var list = []int{idx}
	var indent = p.getItemIndent(idx)
	for i := idx + 1; i < len(p); i++ {
		if itemIndent := p.getItemIndent(i); itemIndent == indent {
			list = append(list, i)
		} else if itemIndent < indent {
			break
		}
	}
	return list
}

func (p SidebarTree) Children(idx int) []int {
	if idx >= len(p)-1 {
		return nil
	}
	if p.getItemIndent(idx) < p.getItemIndent(idx+1) {
		return p.Siblings(idx + 1)
	}
	return nil
}

func (p SidebarTree) getItemIndent(idx int) int {
	for i, ch := range p[idx].Prefix {
		if ch != ' ' {
			return i
		}
	}
	return 0
}
