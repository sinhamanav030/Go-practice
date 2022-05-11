package wedevtoolkit

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func HmacTut() {
	fmt.Println(getCode("sample-text"))
	fmt.Println(getCode("fdsczx"))
}

func getCode(s string) string {
	h := hmac.New(sha256.New, []byte("gophers"))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func EncodingTut() {
	s := "hjgcxbzhjcbdshbxngjsagdvgsacdsajhgdxmjdxnmas723i102bfcu9qwsjcdsi9whdjndsiqwrhdhiw9q"

	encodestd := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+"
	fmt.Println(len(encodestd))
	s64 := base64.NewEncoding(encodestd).EncodeToString([]byte(s))
	s64b := base64.StdEncoding.EncodeToString([]byte(s))
	fmt.Println(s64)
	fmt.Println(s64b)

	// dec64, err := base64.NewEncoding(encodestd).DecodeString(s64)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(dec64)

	dec64b, err := base64.StdEncoding.DecodeString(s64b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(dec64b))

}

func ContextTut() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, "USERID", 777)
	ctx = context.WithValue(ctx, "fname", "bond")

	results, err := dbAccess(ctx)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	fmt.Fprint(w, results)
}

func dbAccess(ctx context.Context) (int, error) {

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	ch := make(chan int)
	go func() {
		uid := ctx.Value("USERID").(int)
		time.Sleep(10 * time.Second)
		if ctx.Err() != nil {
			return
		}
		ch <- uid
	}()

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case i := <-ch:
		return i, nil
	}
}

func bar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Fprint(w, ctx)
}

func HttpsTut() {
	http.HandleFunc("/", foo)
	http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)
}
