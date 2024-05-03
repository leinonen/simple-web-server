package main

import (
	"html/template"
	"net/http"
)

type ViewData struct {
	Message string
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("index").Parse(`
       <!DOCTYPE html>
       <html>
       <head>
           <title>Simple Web Server</title>
       </head>
       <body>
           <h1>{{.Message}}</h1>
       </body>
       </html>
       `))

	data := ViewData{Message: "Hello, World!"}
	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", mainHandler)
	http.ListenAndServe(":8080", nil)
}
