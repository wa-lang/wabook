// Copyright 2024 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mnbook

import (
	"io/fs"
	"path/filepath"
	"strings"
)

// 加载 *.talk.md 文件列表
func loadTalks(root string) []string {
	var ss []string
	filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if strings.HasSuffix(strings.ToLower(path), ".talk.md") {
			ss = append(ss, path)
		}
		return nil
	})
	return ss
}