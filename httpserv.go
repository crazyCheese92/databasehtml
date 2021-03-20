package main

import (
	"net/http"
	"io"
	"fmt"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		data, err := io.ReadAll(req.Body)
		req.Body.Close()
		if err != nil {return }
		
		fmt.Printf("%s\n", data)
		io.WriteString(w, "successful post")
	} else {
		w.WriteHeader(405)
	}
}

func main() {
	http.HandleFunc("/", Handler)
	
	err := http.ListenAndServe(":8080", nil)
	panic(err)
}
