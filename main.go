package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
)

const buffSize = 8

func main() {
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("New connection from", conn.RemoteAddr())
		for line := range getLinesChannel(conn) {
			fmt.Println(line)
		}
		fmt.Println("Closing Connection")
	}

}

func getLinesChannel(f io.ReadCloser) <-chan string {
	buffer := make([]byte, buffSize)
	ch := make(chan string)
	go func() {
		readLine := ""
		for {
			// when fd.read, it stores in buffer and return how many byte read (int)
			bytesRead, err := f.Read(buffer)
			if bytesRead > 0 {
				parts := bytes.Split(buffer[:bytesRead], []byte("\n"))
				readLine += string(parts[0])
				if len(parts) > 1 {
					for _, part := range parts[1:] {
						// found next line, sending current readLine buffer to channel and clear readLine.
						ch <- readLine
						// hold the next part
						readLine = string(part)
					}
				}
			}
			if errors.Is(err, io.EOF) {
				// check if readline was print, if not print it out
				if len(readLine) > 0 {
					ch <- readLine
				}
				f.Close()
				close(ch)
				break
			}
			if err != nil {
				// better to return out to func caller but keep it for now for scope of tutorial
				log.Fatal(err)
			}
		}
	}()
	return ch

}
