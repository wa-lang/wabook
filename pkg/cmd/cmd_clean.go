// Copyright 2023 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	cli "github.com/urfave/cli/v2"
)

var CmdClean = &cli.Command{
	Name:      "clean",
	Usage:     "Deletes a built book",
	ArgsUsage: "[dir]",
	Action: func(ctx *cli.Context) error {
		dir := ctx.Args().First()
		if dir == "" {
			dir, _ = os.Getwd()
		}
		if err := CleanBook(dir); err != nil {
			fmt.Println("delete book failed:", err)
			os.Exit(1)
		}
		return nil
	},
}

func CleanBook(path string) error {
	return os.RemoveAll(filepath.Join(path, "book"))
}
