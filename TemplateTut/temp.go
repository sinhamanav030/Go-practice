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
	tpl, err := tpl.ParseFiles("./TemplateTut/one.txt")
	fruits := []string{"orange", "mango", "banana"}
	err = tpl.ExecuteTemplate(os.Stdout, "two.txt", fruits)
	if err != nil {
		log.Fatalln(err)
	}

}

type course struct {
	Number string
	Name   string
	Units  string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	AcaYear string
	Fall    semester
	Spring  semester
	Summer  semester
}

func Task1() {

	tpl := template.Must(template.ParseFiles("./TemplateTut/Files/task1.gohtml"))
	years := []year{
		year{
			AcaYear: "2020-2021",
			Fall: semester{
				Term: "Fall",
				Courses: []course{
					course{"CSCI-40", "Introduction to Programming in Go", "4"},
					course{"CSCI-130", "Introduction to Web Programming with Go", "4"},
					course{"CSCI-140", "Mobile Apps Using Go", "4"},
				},
			},
			Spring: semester{
				Term: "Spring",
				Courses: []course{
					course{"CSCI-50", "Advanced Go", "5"},
					course{"CSCI-190", "Advanced Web Programming with Go", "5"},
					course{"CSCI-191", "Advanced Mobile Apps With Go", "5"},
				},
			},
		},
		year{
			AcaYear: "2021-2022",
			Fall: semester{
				Term: "Fall",
				Courses: []course{
					course{"CSCI-40", "Introduction to Programming in Go", "4"},
					course{"CSCI-130", "Introduction to Web Programming with Go", "4"},
					course{"CSCI-140", "Mobile Apps Using Go", "4"},
				},
			},
			Spring: semester{
				Term: "Spring",
				Courses: []course{
					course{"CSCI-50", "Advanced Go", "5"},
					course{"CSCI-190", "Advanced Web Programming with Go", "5"},
					course{"CSCI-191", "Advanced Mobile Apps With Go", "5"},
				},
			},
		},
	}

	task1, err := os.Create("./TemplateTut/Output/task1.html")

	err = tpl.Execute(task1, years)
	if err != nil {
		log.Fatalln(err)
	}

}
