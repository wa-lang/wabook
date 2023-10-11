// Copyright 2023 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mnbook

import (
	"reflect"
	"testing"
)

func TestLoadSummaryFrom(t *testing.T) {
	got, err := LoadSummaryFrom("SUMMARY.md", []byte(tSummaryData))
	if err != nil {
		t.Fatal(err)
	}

	expect := tExpectSummary
	if !reflect.DeepEqual(got, expect) {
		t.Fatalf("\ngot: %v\nexpect: %v", got, expect)
	}
}

var tExpectSummary = &Summary{
	Title: "Summary",
	Chapters: []SummaryItem{
		{Name: "Book", Location: "index.md"},
		{Name: "Preface", Location: "preface.md"},

		{Prefix: "-", Name: "Chapter 1", Location: "./src/chapter_1.md"},
		{Prefix: "  -", Name: "Chapter 1.1", Location: "./src/chapter_1.1.md"},
		{Prefix: "  -", Name: "Chapter 1.2", Location: "./src/chapter_1.2.md"},
		{Prefix: "-", Name: "Chapter 2", Location: "./src/chapter_2.md"},
	},
}

const tSummaryData = `
<!-- comment -->

# Summary

[Book](index.md)
[Preface](preface.md)

- [Chapter 1](./src/chapter_1.md)
  - [Chapter 1.1](./src/chapter_1.1.md)
  - [Chapter 1.2](./src/chapter_1.2.md)

- [Chapter 2](./src/chapter_2.md)

<!-- comment -->
`
