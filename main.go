// Copyright 2019 Pierre-Yves Ritschard <pyr@spootnik.org>

// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// notice and this permission notice appear in all copies.

// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

package main

import (
	"bufio"
	"fmt"
	"github.com/tj/go-termd"
	"io/ioutil"
	"os"
)

func main() {

	var c termd.Compiler

	c.SyntaxHighlighter = termd.SyntaxTheme{
		"keyword": termd.Style{},
		"comment": termd.Style{
			Color: "#323232",
		},
		"literal": termd.Style{
			Color: "#555555",
		},
		"name": termd.Style{
			Color: "#777777",
		},
		"name.function": termd.Style{
			Color: "#444444",
		},
		"literal.string": termd.Style{
			Color: "#333333",
		},
	}

	if len(os.Args) == 1 {
		contents, err := ioutil.ReadAll(bufio.NewReader(os.Stdin))
		if err != nil {
			fmt.Fprintln(os.Stderr, "cannot read from stdin:", err)
			os.Exit(1)
		}
		fmt.Println(c.Compile(string(contents)))
	}
	for _, path := range os.Args[1:] {
		contents, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Fprintln(os.Stderr, "cannot open file:", err)
			os.Exit(1)
		}
		fmt.Println(c.Compile(string(contents)))
	}
}
