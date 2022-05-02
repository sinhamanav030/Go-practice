package statesession

import (
	// "fmt"
	// "io"
	"log"
	"net/http"

	// "os"
	// "path/filepath"
	"text/template"
)

func UrlPassTut() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		v := r.FormValue("fname")
		tpl := template.Must(template.New("form.html").ParseFiles("./StateSession/form.html"))
		err := tpl.Execute(w, v)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Fprintf(w, "search for %s", v)
	})

	http.ListenAndServe(":8080", nil)
}

func FileHandleTut() {

	http.HandleFunc("/", fileHandler)

	http.ListenAndServe(":8080", nil)
}

func fileHandler(w http.ResponseWriter, r *http.Request) {

	// var s string
	bs := make([]byte, r.ContentLength)
	r.Body.Read(bs)
	// if r.Method == http.MethodPost {

	// 	f, h, err := r.FormFile("file")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	defer f.Close()
	// 	fmt.Println(f, h, err)
	// 	data, err := io.ReadAll(f)

	// 	if err != nil {
	// 		log.Fatal("error reading file")
	// 	}

	// 	dst, err := os.Create(filepath.Join("./", h.Filename))
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	_, err = dst.Write(data)
	// 	if err != nil {
	// 		http.Error(w, "Something Went Wrong", http.StatusInternalServerError)
	// 		log.Fatal(err)
	// 	}
	// 	s = string(data)
	// }

	tpl := template.Must(template.New("fileform.html").ParseFiles("./StateSession/fileform.html"))
	// fmt.Println(s)
	err := tpl.Execute(w, string(bs))

	if err != nil {
		log.Fatal(err)
	}

}

func redirectTut() {

}
