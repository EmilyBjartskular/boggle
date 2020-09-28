package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Reads all .txt files in the current folder
// and encodes them as strings literals in textfiles.go
func main() {
	fs, err := ioutil.ReadDir(".")
	if err != nil {
		panic(err)
	}
	file, err := os.Create("textfiles.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	w := bufio.NewWriter(file)

	file.Write([]byte("package main \n\nvar (\n"))

	for _, f := range fs {
		if strings.HasSuffix(f.Name(), ".txt") {
			fmt.Fprintf(w, "  %s = []string {\n", strings.TrimSuffix(f.Name(), ".txt"))
			in, _ := os.Open(f.Name())
			out := bufio.NewScanner(in)
			for out.Scan() {
				fmt.Fprintf(w, "    \"%s\",\n", out.Text())
			}
			fmt.Fprint(w, "  }\n")
		}
	}
	fmt.Fprint(w, ")\n")
	w.Flush()
}
