package main

import (
	"bytes"
	"flag"
	"fmt"
	"html"
	"log"
	"os"
	"text/template"
	"time"
)

var holder string

func init() {
	flag.StringVar(&holder, "holder", "Georg Held", "the Copyright holder")
}

type templInfo struct {
	Year   int
	Holder string
}

func main() {
	liFile, err := os.Create("LICENSE.md")
	if err != nil {
		log.Fatal("Creating LICENSE.md: ", err)
	}
	defer liFile.Close()
	reFile, err := os.Create("README.md")
	if err != nil {
		log.Fatal("Creating README.md: ", err)
	}
	defer reFile.Close()

	info := templInfo{Year: time.Now().Year(), Holder: holder}
	templ := template.Must(template.New("name").Parse(templateString))
	var buffer bytes.Buffer
	templ.Execute(&buffer, info)
	s := html.UnescapeString(buffer.String())
	fmt.Fprint(liFile, s)
	fmt.Fprint(reFile, s)
}
