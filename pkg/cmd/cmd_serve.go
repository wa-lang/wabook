// Copyright 2023 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	cli "github.com/urfave/cli/v2"
)

var CmdServe = &cli.Command{
	Name:      "serve",
	Usage:     "Serves a book at http://localhost:3000",
	ArgsUsage: "[dir]",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "http",
			Usage: "HTTP service address",
			Value: "localhost:3000",
		},
	},
	Action: func(ctx *cli.Context) error {
		dir := ctx.Args().First()
		if dir == "" {
			dir, _ = os.Getwd()
		}

		Serve(dir, ctx.String("http"))
		return nil
	},
}

func Serve(path, addr string) {
	fmt.Println("====== Build Book ======")
	if err := BuildBook(path); err != nil {
		log.Fatal("Build failed: ", err)
	}

	fmt.Println("====== Start Http server ======")
	switch {
	case strings.HasPrefix(addr, ":"):
		fmt.Println("listen on http://localhost" + addr)
	case strings.HasPrefix(addr, "http"):
		fmt.Println("listen on " + addr)
	default:
		fmt.Println("listen on http://" + addr)
	}

	h := http.FileServer(http.Dir(filepath.Join(path, "book")))
	if err := http.ListenAndServe(addr, h); err != nil {
		log.Fatal("Serve failed: ", err)
	}

	<-make(chan bool)
}
