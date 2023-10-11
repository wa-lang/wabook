// Copyright 2023 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package render

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func cpDir(dst, src string, ignores []string) error {
	entryList, err := os.ReadDir(src)
	if err != nil && !os.IsExist(err) {
		return err
	}

Loop:
	for _, entry := range entryList {
		if strings.HasSuffix(strings.ToLower(entry.Name()), ".md") {
			continue
		}

		for _, s := range ignores {
			if strings.EqualFold(s, entry.Name()) {
				continue Loop
			}
		}

		if entry.IsDir() {
			if err = cpDir(dst+"/"+entry.Name(), src+"/"+entry.Name(), ignores); err != nil {
				return err
			}
		} else {
			srcFname := filepath.Clean(src + "/" + entry.Name())
			dstFname := filepath.Clean(dst + "/" + entry.Name())
			fmt.Printf("copy %s\n", srcFname)

			if err = cpFile(dstFname, srcFname); err != nil {
				return err
			}
		}
	}

	return nil
}

func cpFile(dst, src string) error {
	err := os.MkdirAll(filepath.Dir(dst), 0777)
	if err != nil && !os.IsExist(err) {
		return err
	}

	fsrc, err := os.Open(src)
	if err != nil {
		return err
	}
	defer fsrc.Close()

	fdst, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer fdst.Close()
	if _, err = io.Copy(fdst, fsrc); err != nil {
		return err
	}

	return nil
}
