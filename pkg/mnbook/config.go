// Copyright 2023 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mnbook

import (
	"bytes"
	"os"

	"github.com/BurntSushi/toml"
)

type BookToml struct {
	Book       BookConfig `json:"book" toml:"book"`
	OutputHtml HtmlConfig `json:"output.html" toml:"output.html"`
}

type BookConfig struct {
	Authors     []string `json:"authors" toml:"authors"`
	Description string   `json:"description" toml:"description"`
	Language    string   `json:"language" toml:"language"`
	Src         string   `json:"src" toml:"src"`
	Title       string   `json:"title" toml:"title"`
}

type HtmlConfig struct {
	GitRepositoryIcon string `json:"git-repository-icon" toml:"git-repository-icon"`
	GitRepositoryUrl  string `json:"git-repository-url" toml:"git-repository-url"`
	EditUrlTemplate   string `json:"edit-url-template" toml:"edit-url-template"`
}

func LoadConfig(path string) (meta *BookToml, err error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return
	}

	data = bytes.Replace(data, []byte("[output.html]"), []byte(`["output.html"]`), 1)
	meta = new(BookToml)
	_, err = toml.Decode(string(data), meta)
	return
}

func (p *BookToml) String() string {
	var buf bytes.Buffer
	enc := toml.NewEncoder(&buf)
	enc.Indent = ""

	err := enc.Encode(p)
	if err != nil {
		panic(err)
	}

	data := bytes.Replace(buf.Bytes(), []byte(`["output.html"]`), []byte("[output.html]"), 1)
	return string(data)
}
