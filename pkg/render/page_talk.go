// Copyright 2024 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package render

import (
	"bytes"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/wa-lang/mnbook/pkg/mntalk/present"
)

func (p *BookRendor) renderAllTalkPages() error {
	if len(p.Book.Talks) == 0 {
		return nil
	}

	for _, s := range p.Book.Talks {
		if err := p.renderTalkPages(s); err != nil {
			return err
		}
	}

	return nil
}

func (p *BookRendor) renderTalkPages(name string) error {
	content, err := os.ReadFile(name)
	if err != nil {
		return err
	}

	doc, err := present.Parse(bytes.NewReader(content), name, 0)
	if err != nil {
		return err
	}

	var fnMap = template.FuncMap{}
	t := template.Must(template.New("").Funcs(fnMap).Parse(tmplTalk))

	var buf bytes.Buffer
	if err = doc.Render(&buf, t); err != nil {
		return err
	}

	relpath := name
	dstAbsPath := filepath.Join(p.Book.Root, "book", relpath)
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
