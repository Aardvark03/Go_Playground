package main

import (
    "fmt"
    "net/http"
)

func handler(writer http.ResponseWriter, reader *http.Request) {
    fmt.Fprintf(writer, "<html><body><a href='/Users/ole/Documents/Go/hello.txt' download>Download File</a></body></html>") 
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8088", nil)
}
