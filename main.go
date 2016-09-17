package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"unicode/utf8"
)

func main() {
	fileName := flag.String("fileName", "", "file name")
	flag.Parse()

	data, err := ioutil.ReadFile(*fileName)
	if err != nil {
		log.Fatal(err)
	}

	trimed := bytes.TrimRight(data, "\n")
	result := utf8.RuneCount(trimed)

	fmt.Printf("This file has %d characters\n", result)
}
