package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var isNumbered = flag.Bool("n", false, "print line number")
var lineNumber = 1

func main() {
	flag.Parse()

	args := flag.Args()

	for _, arg := range args {
		printFile(arg)
	}
}

func printFile(fileName string) {
	fp, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
