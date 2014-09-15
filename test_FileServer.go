package main

import (
	"fmt"
	"net/http"
)

func OurLoggingHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(*r.URL)
		h.ServeHTTP(w, r)
	})
}

func main() {
	fileHandler := http.FileServer(http.Dir("/Users/yinxiaoliu/web/go/src/appmanager/apk"))
	wrappedHandler := OurLoggingHandler(fileHandler)
	http.ListenAndServe(":8080", wrappedHandler)
}
