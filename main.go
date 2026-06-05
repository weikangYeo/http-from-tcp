package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fd, err := os.Open("message.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	// create a 8 Byte buffer to read from file base on requirement
	buffer := make([]byte, 8)
	readLine := ""
	for {
		// when fd.read, it stores in buffer and return how many byte read (int)
		byteRead, err := fd.Read(buffer)
		if byteRead > 0 {
			parts := bytes.Split(buffer[:byteRead], []byte("\n"))
			readLine += string(parts[0])
			if len(parts) > 1 {
				for _, part := range parts[1:] {
					// contain next line, print out and clear readLine
					fmt.Printf("read: %s\n", readLine)
					// hold the next part
					readLine = string(part)
				}
			}
		}
		if err == io.EOF {
			// check if readline was print, if not print it out
			if len(readLine) > 0 {
				fmt.Printf("read: %s\n", readLine)
			}
			break
		}
		if err != nil {
			log.Fatal(err)
		}
	}

}
