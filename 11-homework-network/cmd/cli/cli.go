package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to searching server established!")
	defer conn.Close()

	for {
		sreq := ""
		_, err := fmt.Scanln(&sreq)
		if err != nil {
			fmt.Println(err)
			continue
		}

		_, err = conn.Write([]byte(sreq + "\n"))
		if err != nil {
			fmt.Println(err)
			continue
		}

		conn.SetReadDeadline(time.Now().Add(time.Second * 6))
		time.Sleep(time.Second)

		var sres []byte
		sres, err = io.ReadAll(conn)
		if err != nil && !errors.Is(err, os.ErrDeadlineExceeded) {
			fmt.Println(err)
			continue
		}
		fmt.Printf("%s", sres)
	}

}
