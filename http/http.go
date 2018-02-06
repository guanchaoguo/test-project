package main

import (
	"io"
	"net/http"
)


func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w,req.RemoteAddr )
}

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":12345", nil)

}