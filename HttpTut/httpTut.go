package Httptut

import (
	// "fmt"
	// "log"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"text/template"
)

type myhandler int

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./HttpTut/Templates/*.html"))
}

func (m myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "I am her in handler")
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	data := struct {
		Method        string
		Url           *url.URL
		Submissions   map[string][]string
		Header        http.Header
		Host          string
		ContentLength int64
	}{
		r.Method,
		r.URL,
		r.Form,
		r.Header,
		r.Host,
		r.ContentLength,
	}
	tpl.ExecuteTemplate(w, "req-temp.html", data)

}

func HttpServer() {
	var h myhandler
	http.ListenAndServe(":8080", h)
}

func HelloWorldHttpServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("You have requested", r.URL.Path)
	})
	http.ListenAndServe(":8080", nil)
}

func HttpServerTut() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to my http server", r.URL.Query().Get("token"))
	})
	fs := http.FileServer(http.Dir("static/"))

	http.Handle("/static/", http.StripPrefix("static/", fs))
	http.ListenAndServe(":8080", nil)
}

func HttpServeFileTut() {

	http.HandleFunc("/", setTemp)
	// http.HandleFunc("/img.webp", setImg)
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./HttpTut/static"))))

	http.ListenAndServe(":8080", nil)
}

func setImg(w http.ResponseWriter, req *http.Request) {
	// f, err := os.Open("./HttpTut/static/img.webp")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer f.Close()

	// // io.Copy(w, f)

	// fi, err := f.Stat()
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// http.ServeContent(w, req, fi.Name(), fi.ModTime(), f)

	http.ServeFile(w, req, "./HttpTut/static/img.webp")
}

func setTemp(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `<img src="static/img.webp">`)
}
