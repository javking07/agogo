package main

import (
	"fmt"
	"text/template"
)

var tpl *template.Template

func main() {
	tpl.ParseFiles("sample.gohtml")
	tpl
	fmt.Println(tpl)
}
