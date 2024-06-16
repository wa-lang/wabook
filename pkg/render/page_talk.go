// Copyright 2024 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package render

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
	return nil
}
