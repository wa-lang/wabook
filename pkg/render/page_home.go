// Copyright 2023 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package render

import (
	"os"
	"path/filepath"
)

func (p *BookRendor) renderHomepage() error {
	return os.WriteFile(filepath.Join(p.Book.Root, "book", "index.html"), []byte("hello mnbook"), 0666)
}
