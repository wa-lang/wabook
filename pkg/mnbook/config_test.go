// Copyright 2023 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mnbook

import (
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	meta, err := LoadConfig("../../testdata/book.toml")
	if err != nil {
		t.Fatal(err)
	}
	if got, expect := *meta, tExpect; !reflect.DeepEqual(got, expect) {
		t.Fatalf("\ngot: %v\nexpect: %v", got, expect)
	}
}

func TestBookToml_String(t *testing.T) {
	data, err := os.ReadFile("../../testdata/book.toml")
	if err != nil {
		t.Fatal(err)
	}

	expect := strings.TrimSpace(string(data))
	got := strings.TrimSpace(tExpect.String())

	if got != expect {
		t.Fatalf("got:\n%s\n---\nexpect:\n%s\n", got, expect)
	}
}

var tExpect = BookToml{
	Book: BookConfig{
		Title:    "Book Title",
		Authors:  []string{"chai2010"},
		Language: "zh",
		Src:      ".",
	},
	OutputHtml: HtmlConfig{
		GitRepositoryIcon: "fa-github",
		GitRepositoryUrl:  "https://github.com/wa-lang/mnbook",
		EditUrlTemplate:   "https://github.com/wa-lang/mnbook/edit/master/testdata/{path}",
	},
}
