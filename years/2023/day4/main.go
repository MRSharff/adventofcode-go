package main

import (
	_ "embed"
	"io"
	"strings"
)

//go:embed input.txt
var input []byte

func testInput() io.Reader {
	s := `
19, 13, 30 @ -2,  1, -2
18, 19, 22 @ -1, -1, -2
20, 25, 34 @ -2, -2, -4
12, 31, 28 @ -1, -2, -1
20, 19, 15 @  1, -5, -3`

	r := strings.NewReader(s)
	r.Read([]byte{'\n'})
	return r
}

func main() {

}

func part1(r io.Reader) int {
	t := 0
	a := 19-2*t == 18-1*t

}
