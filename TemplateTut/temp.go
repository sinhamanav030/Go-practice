package Templatetut

import (
	"log"
	"os"
	"strings"
	"text/template"
)

type student struct {
	Name string
	Sno  int
}

func TemplateTut() {
	tpl, err := template.ParseFiles("./TemplateTut/hello.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	// // err = tpl.Execute(os.Stdout, nil)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// nf, err := os.Create("Index.html")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// err = tpl.Execute(nf, nil)

	tpl, err = tpl.ParseFiles("./TemplateTut/one.txt", "./TemplateTut/two.txt")

	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "one.txt", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.txt", nil)

	if err != nil {
		log.Fatalln(err)
	}

	temp := template.Must(template.ParseGlob("TemplateTut/*"))

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	fruits := []string{"orange", "mango", "banana"}

	err = temp.ExecuteTemplate(os.Stdout, "hello.gohtml", fruits)
	if err != nil {
		log.Fatalln(err)
	}

	fruit := make(map[string]string)

	fruit["orange"] = "orange"
	fruit["banana"] = "yellow"
	fruit["peach"] = "peach"

	err = temp.ExecuteTemplate(os.Stdout, "hello.gohtml", fruit)
	if err != nil {
		log.Fatalln(err)
	}

	s1 := student{"Manav", 1}
	s2 := student{"karan", 2}
	s3 := student{"Rasmi", 3}

	s := []student{s1, s2, s3}

	err = temp.ExecuteTemplate(os.Stdout, "one.txt", s)

}

func FuncBindTut() {
	fm := template.FuncMap{
		"uc": strings.ToUpper,
		"lc": strings.ToLower,
	}

	tpl := template.Must(template.New("").Funcs(fm).ParseFiles("./TemplateTut/two.txt"))

	fruits := []string{"orange", "mango", "banana"}
	err := tpl.ExecuteTemplate(os.Stdout, "two.txt", fruits)
	if err != nil {
		log.Fatalln(err)
	}
}
