# Copyright 2023 <chaishushan{AT}gmail.com>. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

defaule:
	go test ./...
	go run main.go

build:
	go run main.go build ./testdata

serve:
	go run main.go serve ./testdata

debug:
	-rm -rf ./testdata/book
	-rm -rf ./testdata/_book

	go run main.go build ./testdata && cd ./testdata && mv book _book_go
	cd ./testdata && wabook build && mv book _book
	cd ./testdata && mv _book_go book

clean:
	-go run main.go clean ./testdata
