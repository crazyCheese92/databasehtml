package main

import (
	"fmt"
	"net/http"
	"io"
	"io/ioutil"
	"bytes"
)

func main() {
	client := &http.Client{}
	
	
	var bodyWrite string
	fmt.Print("Write: ")
	fmt.Scan(&bodyWrite)
	var body bytes.Buffer
	body.Write([]byte("bodyWrite"))
	
	var request string
	fmt.Print("Request (POST/PUT): ")
	fmt.Scan(&request)
	req, err := http.NewRequest("request", "http://localhost:8080/", &body)
	
	resp, err := client.Do(req)
	if err != nil {panic(err) }
	
	data, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	
	err = ioutil.WriteFile("p.html", data, 0666)
	if err != nil {panic(err) }
	
	
}
