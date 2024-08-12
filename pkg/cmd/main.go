// Copyright 2023 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"os"
	"runtime/debug"
	"strings"

	cli "github.com/urfave/cli/v2"
)

func Main() {
	cliApp := cli.NewApp()
	cliApp.Name = "wabook"
	cliApp.Usage = "A tool for build mini markdown book"
	cliApp.HideHelpCommand = true
	cliApp.Version = func() string {
		if info, ok := debug.ReadBuildInfo(); ok {
			if s := info.Main.Version; s != "" {
				return s
			}
		}
		return "dev"
	}()

	cliApp.CustomAppHelpTemplate = cli.AppHelpTemplate +
		"\n See \"https://github.com/wa-lang/wabook\" for more information.\n"

	cliApp.Action = func(ctx *cli.Context) error {
		if ctx.NArg() > 0 {
			fmt.Println("unknown command:", strings.Join(ctx.Args().Slice(), " "))
			os.Exit(1)
		}
		cli.ShowAppHelpAndExit(ctx, 0)
		return nil
	}

	cliApp.Commands = []*cli.Command{
		CmdInit,
		CmdBuild,
		CmdServe,
		CmdClean,
	}

	cliApp.Run(os.Args)
}
