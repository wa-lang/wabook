// Copyright 2023 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package render

import (
	"embed"
	"io/fs"
	"os"
	"path/filepath"
)

//go:embed _static
var _staticFS embed.FS

func (p *BookRendor) renderStaticFile() error {
	return cpDir(filepath.Join(p.Book.Root, "book"), p.Book.Root, []string{
		".git", "_git",
		"book", "book.toml",
		"_book",
		"_book1",
		"_book2",
		"_book3",
	})
}

func (p *BookRendor) renderStaticAsset() error {
	staticFS, err := fs.Sub(_staticFS, "_static")
	if err != nil {
		return err
	}
	err = fs.WalkDir(staticFS, ".", func(path string, d fs.DirEntry, err error) error {
		if d == nil || d.IsDir() {
			return nil
		}
		if err != nil {
			return err
		}

		data, err := fs.ReadFile(staticFS, path)
		if err != nil {
			return err
		}

		dstpath := filepath.Join(p.Book.Root, "book", path)
		os.MkdirAll(filepath.Dir(dstpath), 0777)

		os.WriteFile(dstpath, data, 0666)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
