package main

import (
	"fmt"
	"net/http"
)

func main()  {



	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("ye")

	})

	http.ListenAndServe(":9999", nil)
}
