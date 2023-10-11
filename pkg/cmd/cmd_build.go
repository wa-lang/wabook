// Copyright 2023 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"os"

	cli "github.com/urfave/cli/v2"

	"github.com/wa-lang/mnbook/pkg/mnbook"
	"github.com/wa-lang/mnbook/pkg/render"
)

var CmdBuild = &cli.Command{
	Name:      "build",
	Usage:     "Builds a book from its markdown files",
	ArgsUsage: "[dir]",
	Action: func(ctx *cli.Context) error {
		dir := ctx.Args().First()
		if dir == "" {
			dir, _ = os.Getwd()
		}
		if err := BuildBook(dir); err != nil {
			fmt.Println("build book failed:", err)
			os.Exit(1)
		}
		return nil
	},
}

func BuildBook(path string) error {
	book, err := mnbook.LoadBook(path)
	if err != nil {
		return fmt.Errorf("LoadBook: %w", err)
	}

	if err := render.New().Run(book); err != nil {
		return err
	}

	return nil
}
