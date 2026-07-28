package main

import (
	"fmt"
	"net/http"

	Router "github.com/sebastian-ruiz-7/http-go-server/internal/http"
)

func main() {

	mux := Router.NewRouter()

	err := http.ListenAndServe(":3000", mux)

	if err != nil {
		fmt.Println("error", err)
	}
}
