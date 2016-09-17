package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"unicode/utf8"
)

func main() {
	var data []byte
	var err error

	if len(os.Args) > 2 {
		fileName := flag.String("fileName", "", "file name")
		flag.Parse()

		data, err = ioutil.ReadFile(*fileName)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		data = []byte(os.Args[1])
	}

	trimed := bytes.TrimRight(data, "\n")
	result := utf8.RuneCount(trimed)

	fmt.Printf("This file has %d characters\n", result)
}
