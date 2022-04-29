package TcpTut

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"log"
	"net"
	// "text/scanner"
)

func Conn() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Panicln(err)
		}
		// io.WriteString(conn, "\nHello\n")
		// fmt.Fprintln(conn, "How")
		// fmt.Fprintf(conn, "%v", "yo")
		// conn.Close()
		go handle(conn)
	}

}

func handle(conn net.Conn) {

	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		fmt.Fprintln(conn, "Connection TimeOut")
	}
	s := bufio.NewScanner(conn)

	for s.Scan() {
		fmt.Println(s.Text())
		fmt.Fprintf(conn, "%v", s.Text())
	}
	defer conn.Close()
	fmt.Println("End")
}

func netDial() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	bs, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Print(bs)
}

func EncConn() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handleRot(conn)
	}

}

func handleRot(conn net.Conn) {
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(conn)

	for s.Scan() {
		ln := strings.ToLower(s.Text())
		bs := []byte(ln)
		r := rot13(bs)
		fmt.Fprintf(conn, "%v\n", string(r))
	}
	defer conn.Close()
	fmt.Println("End")

}

func rot13(b []byte) []byte {
	r := make([]byte, len(b))
	for i, v := range b {
		if b[i] <= 109 {
			r[i] = v + 13
		} else {
			r[i] = v - 13
		}
	}
	return r
}

func HttpServer() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}

		go handleHttp(conn)
	}
}

func handleHttp(conn net.Conn) {
	defer conn.Close()

	request(conn)

	response(conn)
}

func request(conn net.Conn) {
	i := 0
	sc := bufio.NewScanner(conn)
	for sc.Scan() {
		ln := sc.Text()
		fmt.Println(ln)

		if i == 0 {
			m := strings.Fields(ln)[0]
			fmt.Println("***METHOD", m)
		}
		if ln == "" {
			break
		}
		i++
	}
}

func response(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello World</strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
