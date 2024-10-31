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

	"github.com/wa-lang/wabook/pkg/present"
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

func (p *BookRendor) renderTalkPages(path string) error {
	content, err := os.ReadFile(filepath.Join(p.Book.Root, path))
	if err != nil {
		return err
	}

	doc, err := present.Parse(bytes.NewReader(content), path, 0)
	if err != nil {
		return err
	}

	t := template.Must(present.Template().Parse(tmplPresent))
	t = template.Must(t.Parse(tmplAction))

	var htmlContent []byte
	{
		var buf bytes.Buffer
		if err = doc.Render(&buf, t); err != nil {
			return err
		}
		htmlContent = bytes.TrimSpace(buf.Bytes())
	}

	dstAbsPath := filepath.Join(p.Book.Root, "book", path)
	if ext := filepath.Ext(dstAbsPath); strings.EqualFold(ext, ".slide") {
		dstAbsPath = dstAbsPath[:len(dstAbsPath)-len(".slide")]
	}
	dstAbsPath += ".html"

	os.MkdirAll(filepath.Dir(dstAbsPath), 0777)
	if err := os.WriteFile(dstAbsPath, htmlContent, 0666); err != nil {
		return err
	}

	return nil
}
