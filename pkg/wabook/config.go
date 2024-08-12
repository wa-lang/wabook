// Copyright 2023 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wabook

import (
	"bytes"
	"os"

	"github.com/BurntSushi/toml"
)

type BookToml struct {
	Book       BookConfig   `json:"book" toml:"book"`
	Custom     CustomConfig `json:"custom" toml:"custom"`
	Giscus     GiscusConfig `json:"giscus" toml:"giscus"`
	OutputHtml HtmlConfig   `json:"output.html" toml:"output.html"`
}

type CustomConfig struct {
	ContentHeader string `json:"content_header" toml:"content_header"`
	ContentFooter string `json:"content_footer" toml:"content_footer"`
	PageFooter    string `json:"page_footer" toml:"page_footer"`
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

type GiscusConfig struct {
	Enabled        bool   `json:"enaled" toml:"enaled"`
	DataRepo       string `json:"data_repo" toml:"data_repo"`
	DataRepoId     string `json:"data_repo_id" toml:"data_repo_id"`
	DataCategory   string `json:"data_category" toml:"data_category"`
	DataCategoryId string `json:"data_category_id" toml:"data_category_id"`
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
