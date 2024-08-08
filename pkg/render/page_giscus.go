// Copyright 2024 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package render

import (
	"bytes"
	_ "embed"
	"text/template"
)

//go:embed tmpl/giscus.js
var tmplGenGiscusJs string

func (p *BookRendor) genGiscusJs() (string, error) {
	var buf bytes.Buffer

	t := template.Must(template.New("giscus.js").Parse(tmplGenGiscusJs))
	err := t.Execute(&buf, p.Book.Info.Giscus)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
