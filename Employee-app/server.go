package EmployeeApp

import "net/http"

func HttpServer() {
	http.Handle("/", http.StripPrefix("", http.FileServer(http.Dir("./Employee-app/Static"))))
	http.ListenAndServe(":8080", nil)
}
