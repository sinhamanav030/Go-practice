package TcpTut

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"time"

	"io"
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
		io.WriteString(conn, "\nHello\n")
		fmt.Fprintln(conn, "How")
		fmt.Fprintf(conn, "%v", "yo")
		conn.Close()
		// go handle(conn)
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

func main() {
	netDial()
}
