package main

import (
	"io"
	"log"
	"net"
	"time"
)

func Example02() {
	listener, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go func() {
			for {
				_, err := io.WriteString(conn, time.Now().Format("15:05:05\n"))
				if err != nil {
					log.Println(err)
					return
				}

				time.Sleep(time.Second)
			}
		}()
	}
}
