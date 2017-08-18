package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl := template.Must(template.ParseFiles("tpl.gohtml"))
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", "KC Golang!")
	if err != nil {
		log.Fatalln(err)
	}
}
