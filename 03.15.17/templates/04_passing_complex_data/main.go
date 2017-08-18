package main

import (
	"log"
	"os"
	"text/template"
)

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

func main() {
	tpl := template.Must(template.ParseFiles("tpl.gohtml"))

	f := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}

	c := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	cars := []car{f, c}

	data := struct {
		Transport []car
	}{
		cars,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
