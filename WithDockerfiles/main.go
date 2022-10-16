package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	myOs, myArch := runtime.GOOS, runtime.GOARCH
	inContainer := "inside"
	if _, err := os.Lstat("/.dockerenv"); err != nil && os.IsNotExist(err) {
		inContainer = "outside"
	}
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "Hello, %s\n", r.UserAgent())
	_, _ = fmt.Fprintf(w, "%s,%s\n", myOs, myArch)
	_, _ = fmt.Fprintf(w, "%s\n", inContainer)
	_, _ = fmt.Fprintf(w, "%s\n", runtime.Version())
}

func main() {
	http.HandleFunc("/d", homeHandler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Printf("hello %s\n", runtime.Version())
	}
}
