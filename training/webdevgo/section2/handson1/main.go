package main

import "text/template"

/*


   Create a data structure to pass to a template which

   contains information about California hotels including Name, Address, City, Zip, Region
   region can be: Southern, Central, Northern
   can hold an unlimited number of hotels


*/
// holds hotel info
type hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
	Region  string
}

type hotels []hotel

var tpl *template.Template

func init() {
	tpl = template.Must(tpl.ParseGlob("/templates/*"))
}

func main() {
	h1 := hotel{"a", "b", "c", "d", "e"}
	hs := hotels{}
}
