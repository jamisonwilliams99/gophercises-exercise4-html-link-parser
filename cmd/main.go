package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"jamisonwilliams99/gophercises-exercise4-htmlparser/link"
	"strings"
)

func main() {
	filename := flag.String("file", "./html_files/ex1.html", "html file that will be parsed")
	flag.Parse()

	fmt.Printf("Opening %v to be parsed...\n", *filename)

	buf, err := ioutil.ReadFile(*filename)
	if err != nil {
		panic(err)
	}
	r := strings.NewReader(string(buf))
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}
