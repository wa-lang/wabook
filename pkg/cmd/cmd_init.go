// Copyright 2023 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	cli "github.com/urfave/cli/v2"
)

//go:embed _example_book
var _exampleBookFS embed.FS

var CmdInit = &cli.Command{
	Name:      "init",
	Usage:     "Creates a new book",
	ArgsUsage: "[dir]",
	Action: func(ctx *cli.Context) error {
		dir := ctx.Args().First()
		if dir == "" {
			dir, _ = os.Getwd()
		}
		if err := InitBook(dir); err != nil {
			fmt.Println("init book failed:", err)
			os.Exit(1)
		}
		return nil
	},
}

func InitBook(name string) error {
	exampleFS, err := fs.Sub(_exampleBookFS, "_example_book")
	if err != nil {
		return err
	}
	err = fs.WalkDir(exampleFS, ".", func(path string, d fs.DirEntry, err error) error {
		if d == nil || d.IsDir() {
			return nil
		}
		if err != nil {
			return err
		}

		data, err := fs.ReadFile(exampleFS, path)
		if err != nil {
			return err
		}

		dstpath := filepath.Join(name, path)
		os.MkdirAll(filepath.Dir(dstpath), 0777)

		os.WriteFile(dstpath, data, 0666)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
