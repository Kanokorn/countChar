package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"unicode/utf8"
)

func main() {
	fileName := os.Args[1]
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	trimed := bytes.TrimRight(data, "\n")
	result := utf8.RuneCount(trimed)

	fmt.Printf("This file has %d characters\n", result)
}
