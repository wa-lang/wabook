// Copyright 2023 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mnbook

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Summary struct {
	Title    string // Summary
	Chapters []SummaryItem
}

type SummaryItem struct {
	Prefix   string // "", "-", "*", "  -", "  *"
	Name     string
	Location string
}

func LoadSummary(path string) (*Summary, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return newSummaryParser(string(data)).Parse()
}

func LoadSummaryFrom(path string, data []byte) (*Summary, error) {
	return newSummaryParser(string(data)).Parse()
}

func (p *Summary) String() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "# %s\n\n", p.Title)

	for _, item := range p.Chapters {
		fmt.Fprintf(&buf, "%s [%s](%s)\n", item.Prefix, item.Name, item.Location)
	}

	fmt.Fprintln(&buf)
	return buf.String()
}

type summaryParser struct {
	content string
	pos     int
}

func newSummaryParser(s string) *summaryParser {
	s = strings.ReplaceAll(s, "\r\n", "\n")
	return &summaryParser{content: s}
}

func (p *summaryParser) Parse() (*Summary, error) {
	var summary Summary

	for !p.atEnd() {
		line := p.readLine()
		switch {
		case strings.HasPrefix(line, "#"):
			// # Summary
			summary.Title = p.parseTitle(line)

		case strings.Contains(line, "]("):
			// - [Chapter 1](./src/chapter_1.md)
			summary.Chapters = append(summary.Chapters, p.parseItem(line))

		default:
			if s := strings.TrimSpace(line); s != "" {
				return nil, errors.New("summary: invalid item: " + s)
			}
		}
	}

	return &summary, nil
}

func (p *summaryParser) atEnd() bool {
	return p.pos >= len(p.content)
}

func (p *summaryParser) skipComment() {
	if p.atEnd() {
		return
	}
	if !strings.HasPrefix(p.content[p.pos:], "<!--") {
		return
	}
	if end := strings.Index(p.content[p.pos:], "-->"); end > 0 {
		p.pos += end + len("-->")
	} else {
		p.pos = len(p.content)
	}
}

func (p *summaryParser) readLine() string {
	if p.atEnd() {
		return ""
	}
	p.skipComment()

	var line string
	if end := strings.Index(p.content[p.pos:], "\n"); end >= 0 {
		line = p.content[p.pos:][:end]
		if j := strings.Index(line, "<!--"); j >= 0 {
			end = j
		}
		p.pos += end + len("\n")

	} else {
		p.pos = len(p.content)
	}

	return line
}

func (*summaryParser) parseTitle(line string) string {
	return strings.TrimLeft(line, "# ")
}

func (*summaryParser) parseItem(line string) (item SummaryItem) {
	if s := strings.TrimSpace(line); s != "" {
		if s[0] == '-' || s[0] == '*' {
			if i := strings.IndexAny(line, "-*"); i >= 0 {
				item.Prefix = line[:i+1]
				line = strings.TrimSpace(line[i+1:])
			}
		}
	}
	// [text]
	if i := strings.Index(line, "["); i >= 0 {
		line = line[i+1:]
		if j := strings.Index(line, "]"); i >= 0 {
			item.Name = strings.TrimSpace(line[i:j])
			line = line[j+1:]
		}
	}
	// (link)
	if i := strings.Index(line, "("); i >= 0 {
		line = line[i+1:]
		if j := strings.Index(line, ")"); i >= 0 {
			item.Location = strings.TrimSpace(line[i:j])
		}
	}
	return
}
