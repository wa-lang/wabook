// Copyright 2024 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wabook

import (
	"io/fs"
	"path/filepath"
	"strings"
)

// 加载 *.slide 文件列表
func loadTalks(root string) []string {
	var ss []string
	filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if strings.HasSuffix(strings.ToLower(path), ".slide") {
			relpath, err := filepath.Rel(root, path)
			if err == nil {
				ss = append(ss, relpath)
			}
		}
		return nil
	})
	return ss
}
