// Copyright 2023 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package render

import (
	"os"
	"path/filepath"
)

func (p *BookRendor) render404Page() error {
	return os.WriteFile(filepath.Join(p.Book.Root, "book", "404.html"), []byte("TODO"), 0666)
}
