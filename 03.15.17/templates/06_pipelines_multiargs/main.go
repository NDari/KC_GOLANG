package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

func double(x int) int {
	return x + x
}

func square(pow, x int) float64 {
	return math.Pow(float64(x), float64(pow))
}

func sqRoot(x float64) float64 {
	return math.Sqrt(x)
}

func main() {
	fm := template.FuncMap{
		"fdbl":  double,
		"fsq":   square,
		"fsqrt": sqRoot,
	}

	tpl := template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 3)
	if err != nil {
		log.Fatalln(err)
	}
}
