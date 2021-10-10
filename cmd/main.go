package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("Server started at :3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		panic(err)
	}
}
