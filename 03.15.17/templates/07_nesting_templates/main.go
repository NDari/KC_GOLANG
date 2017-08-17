package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl := template.Must(template.ParseGlob("*.gohtml"))
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", 42) // NOTE that we dont need to call defines explicitly
	if err != nil {
		log.Fatalln(err)
	}
}
