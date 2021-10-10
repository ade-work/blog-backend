package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.Handle("/api/hello", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("em1 .. hello")
	}))

	fmt.Println("Server started at :3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		panic(err)
	}
}
