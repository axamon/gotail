// Copyright (c) 2020 Alberto Bregliano

// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:

// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	_ "net/http/pprof"
	"os"
)

var scanner *bufio.Scanner

func main() {

	var body []byte
	var buf = &bytes.Buffer{}

	var num = flag.Int("n", 10, "Number of lines to show")
	flag.Parse()

	input, err := os.Stdin.Stat()
	checkErr(err)

	switch {
	case !(input.Mode()&os.ModeNamedPipe == 0):
		scanner = bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanBytes)
		for scanner.Scan() {
			buf.Write(scanner.Bytes())
		}
		print(buf.Bytes(), *num)
	default:
		for i := range flag.Args() {
			body, err = ioutil.ReadFile(flag.Args()[i])
			checkErr(err)
			if len(flag.Args()) > 1 {
				fmt.Printf("==> %s <==\n", flag.Args()[i])
			}
			print(body, *num)
		}
	}

}

func print(body []byte, num int) {

	scanner = bufio.NewScanner(bytes.NewReader(body))

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for i := len(lines) - num; i < len(lines); i++ {
		fmt.Println(lines[i])
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
