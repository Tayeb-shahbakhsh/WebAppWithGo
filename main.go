package web

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to tayeb's first web application with Go</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Starting the server on port :3000")
	http.ListenAndServe(":3000", nil)
}
