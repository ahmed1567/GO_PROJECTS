package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func readTest() {
	resp ,err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	bs := make([]byte, 2000)

	n, err := resp.Body.Read(bs)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Read", n, "bytes")
	fmt.Println(string(bs[:n]))

}

type logWriter struct{}

func (w logWriter) Write (bs []byte) (int, error){
	fmt.Println(string(bs[:2000]))
	fmt.Println("Number of bytes written:",len(bs))
	return len(bs), nil
}

func writeTest() {
	resp ,err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	w := logWriter{}
	defer resp.Body.Close()

	io.Copy(w, resp.Body)



}
