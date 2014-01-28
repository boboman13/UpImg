package server

import (
	"fmt"
	"net/http"
)

type UpImgServer struct {
	
}

// The index page. This also handles 
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(writer, "<html><head><title>UpImg</title></head><body><h1>UPIMG UPLOAD UR FILEZ</body></html>")
}

// The upload page.
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(writer, "<html><head><title>uploadz</title></head><body>dez uploads</body></html>")
}

// Starts the UpImgServer.
func (*UpImgServer) Start(port string) (output string) {
	http.handleFunc("/", IndexHandler)
	http.handleFunc("/upload", UploadHandler)

	// Listen and serve.
	http.ListenAndServe(port, nil)

	return "Started UpImg on port "+port
}