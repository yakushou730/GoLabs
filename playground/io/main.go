package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://google.com")
	defer resp.Body.Close()
	if err != nil {
		log.Fatalln(err)
		return
	}

	io.Copy(os.Stdout, resp.Body)
}
