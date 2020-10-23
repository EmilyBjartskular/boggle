package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	localeDir = "locale/wl"
)

func main() {
	langs, err := ioutil.ReadDir(localeDir)
	if err != nil {
		panic(err)
	}
	for _, f := range langs {
		if f.IsDir() {
			path := fmt.Sprintf("%s/%s", localeDir, f.Name())
			lists, err := ioutil.ReadDir(path)
			if err != nil {
				panic(err)
			}

			file, err := os.Create(fmt.Sprintf("%s/lists.go", path))
			if err != nil {
				panic(err)
			}
			defer file.Close()

			w := bufio.NewWriter(file)

			fmt.Fprintf(w, "package %s \n\nvar (\n", f.Name())

			for _, f := range lists {
				if strings.HasSuffix(f.Name(), ".txt") {
					filePath := fmt.Sprintf("%s/%s", path, f.Name())
					fmt.Fprintf(w, "  %s = []string {\n", strings.TrimSuffix(f.Name(), ".txt"))
					in, _ := os.Open(filePath)
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
	}
}

// Reads all .txt files in the current folder
// and encodes them as strings literals in textfiles.go
func mainOld() {
	fs, err := ioutil.ReadDir(".")
	if err != nil {
		panic(err)
	}
	file, err := os.Create("list.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	w := bufio.NewWriter(file)

	fmt.Fprint(w, "package main \n\nvar (\n")

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
