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

func (p *BookRendor) tmpl_GenSidebarItems(idx int) string {
	var buf bytes.Buffer
	_ = p.tmpl_genSidebarItems(&buf, idx, 0, 4)
	_ = buf.String()

	return fmt.Sprintf("tmpl_genSidebarItems: idx=%d", idx)
}

func (p *BookRendor) tmpl_genSidebarItems(w io.Writer, idx, level, maxLevel int) string {
	if level >= maxLevel {
		return ""
	}

	indent := strings.Repeat("  ", level)
	if level == 0 {
		fmt.Fprintf(w, indent+`<ol class="chapter">`)
		defer fmt.Fprintf(w, indent+`</ol>`)
	} else {
		fmt.Fprintf(w, indent+`<ol class="section">`)
		defer fmt.Fprintf(w, indent+`</ol>`)
	}

	// TODO
	// <li class="chapter-item expanded ">
	//     <a href="..."><strong aria-hidden="true">1.</strong> Chapter x</a>
	// </li>

	return ""
}
