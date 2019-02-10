package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(w io.Writer, name string) {
	fmt.Fprintf(w, "Porco, %s", name)

}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "Dio!")
}

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}
