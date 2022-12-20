package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	// Your solution goes here. Good luck!

	var hidden = flag.Bool("a", false, "show all")
	flag.Parse()

	fileNames := listFiles("testdata", *hidden)

	for _, fileName := range fileNames {
		fmt.Println(fileName)
	}

}

func listFiles(dirname string, hidden bool) []string {
	var dirs []string

	files, err := ioutil.ReadDir(dirname)

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if !hidden && strings.HasPrefix(f.Name(), ".") {
			continue
		}
		dirs = append(dirs, f.Name())
	}

	return dirs
}
