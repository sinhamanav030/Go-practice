package statesession

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
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

func CookieTut() {
	http.HandleFunc("/", setCookie)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/read", read)
	http.ListenAndServe(":8080", nil)
}

func setCookie(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("counter")
	if err == http.ErrNoCookie {
		http.SetCookie(w, &http.Cookie{
			Name:  "counter",
			Value: "0",
		})
	} else {

		count, err := strconv.Atoi(cookie.Value)
		if err != nil {
			panic(err)
		}
		count++
		cookie.Value = strconv.Itoa(count)
		http.SetCookie(w, cookie)
		fmt.Fprintf(w, "visited :%d", count)
	}

}

func read(w http.ResponseWriter, r *http.Request) {
	cookieData, err := r.Cookie("counter")
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, cookieData)

	// http.SetCookie(w, &http.Cookie{
	// 	Name:  "counter",
	// 	Value: "0",
	// })
}

func UUIDTut() {
	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	u := getUser(r)
	tpl := template.Must(template.ParseFiles("./StateSession/index.gohtml"))
	tpl.Execute(w, u)
}

func getUser(r *http.Request) User {
	var u User
	cookie, err := r.Cookie("session")
	if err != nil {
		return u
	}

	if un, ok := dbSessions[cookie.Value]; ok {
		u = dbUsers[un]
	}

	return u

}

func userFunc(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("user-session")
	if err != nil {
		id := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "user-session",
			Value: id.String(),
			// Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
	}
	fmt.Println(cookie)
}

type user struct {
	UserName string
	First    string
	Last     string
}

type User struct {
	UserName string
	Password []byte
	First    string
	Last     string
}

var dbSessions = map[string]string{}
var dbUsers = map[string]User{}

// func userSession(w http.ResponseWriter, r *http.Request) {
// 	cookie, err := r.Cookie("session")
// 	if err != nil {
// 		sId := uuid.NewV4()
// 		cookie = &http.Cookie{
// 			Name:  "session",
// 			Value: sId.String(),
// 		}
// 		http.SetCookie(w, cookie)
// 	}

// 	var u user

// 	if value, ok := dbSessions[cookie.Value]; ok {
// 		u = dbUsers[value]
// 	}

// 	if r.Method == http.MethodPost {
// 		userName := r.FormValue("username")
// 		f := r.FormValue("firstname")
// 		l := r.FormValue("lastname")
// 		u = user{userName, f, l}
// 		dbSessions[cookie.Value] = userName
// 		dbUsers[userName] = u
// 	}

// 	tpl := template.Must(template.ParseFiles("./StateSession/index.gohtml"))

// 	tpl.Execute(w, u)

// }

func alreadyLoggedIn(req *http.Request) bool {
	cookie, err := req.Cookie("session")
	if err != nil {
		return false
	}
	un := dbSessions[cookie.Value]
	_, ok := dbUsers[un]
	return ok
}

func signup(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		userName := r.FormValue("username")
		pass := r.FormValue("password")
		firstname := r.FormValue("firstname")
		lastname := r.FormValue("lastname")

		if _, ok := dbUsers[userName]; ok {
			http.Error(w, "User Already Exists", http.StatusForbidden)
			return
		}

		sId := uuid.NewV4()

		cookie := &http.Cookie{
			Name:  "session",
			Value: sId.String(),
		}

		http.SetCookie(w, cookie)

		bs, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}

		u := User{userName, bs, firstname, lastname}

		dbSessions[cookie.Value] = userName
		dbUsers[userName] = u

	}
	tpl := template.Must(template.ParseFiles("./StateSession/signup.gohtml"))
	tpl.Execute(w, nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		userName := r.FormValue("username")
		passWord := r.FormValue("password")

		u, ok := dbUsers[userName]
		if !ok {
			http.Error(w, "Username and/or password doesn't match", http.StatusForbidden)
			return
		}
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(passWord))

		if err != nil {
			http.Error(w, "Username and/or password doesn't match", http.StatusForbidden)
			return
		}

		sId := uuid.NewV4()
		cookie := &http.Cookie{
			Name:  "session",
			Value: sId.String(),
		}

		http.SetCookie(w, cookie)

		dbSessions[cookie.Value] = userName
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tpl := template.Must(template.ParseFiles("./StateSession/login.gohtml"))
	tpl.Execute(w, nil)
}

func logout(w http.ResponseWriter, r *http.Request) {
	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	cookie, err := r.Cookie("session")

	if err != nil {
		http.Error(w, "Internal server Error", http.StatusInternalServerError)
		return
	}

	delete(dbSessions, cookie.Value)

	cookie = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}

	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
