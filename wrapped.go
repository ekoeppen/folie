package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func wrappedCd(argv []string) {
	if len(argv) > 1 {
		if err := os.Chdir(argv[1]); err != nil {
			fmt.Println(err)
			return
		}
	}
	if dir, err := os.Getwd(); err == nil {
		fmt.Println(dir)
	} else {
		fmt.Println(err)
	}
}

func wrappedLs(argv []string) {
	dir := "."
	if len(argv) > 1 {
		dir = argv[1]
	}
	var names []string
	files, _ := ioutil.ReadDir(dir)
	for _, f := range files {
		n := f.Name()
		if !strings.HasPrefix(n, ".") {
			if f.IsDir() {
				n = n + "/"
			}
			names = append(names, n)
		}
	}
	fmt.Println(strings.Join(names, " "))
}

func wrappedReset() {
	if dev == nil {
		fmt.Println("[use CTRL-D to exit]")
	} else {
		boardReset(false)
	}
}

func wrappedSend(argv []string) {
	if len(argv) == 1 {
		fmt.Printf("Usage: %s <filename>\n", argv[0])
		return
	}
	if !IncludeFile(argv[1], 0) {
		fmt.Println("Send failed.")
	}
}
