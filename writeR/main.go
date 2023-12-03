package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil{
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// fmt.Println("Response:", resp)

// Reader interface
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	lw := logWriter{}

	io.Copy(lw, resp.Body)

	// or
	// io.Copy(os.Stdout, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	// return 1, nil
	fmt.Println(string(bs))
	fmt.Println("no of bytes:", len(bs))
	return len(bs), nil
}