package main

import (
    "net/http"
	"html/template"
)

type application struct{
	templates *template.Template
}


func main() {
	mux:= http.NewServeMux()

	//Must catches. Almost same as: if err != nil
	var tmpl = template.Must(template.ParseGlob("./ui/html/*.html")) //ParseGLob loads templates IN BULK

	//dependensies 
	app:= &application{
		templates: tmpl,
	}
	

	//if mux recieves "/*" it pulls a specific handle.
    mux.HandleFunc("/", app.HomeHandler)
	mux.HandleFunc("/decoder", app.DecoderHandler)
	mux.HandleFunc("/encoder", app.EncoderHandler)

	//CSS
	assetsDir:= http.Dir("./ui/static")
	fileServer:= http.FileServer(assetsDir)
	mux.Handle("/static/", http.StripPrefix("/static/",fileServer))

	//Creates a server (:port, mux)
    err := http.ListenAndServe(":8082", mux)
	if err != nil{
		panic(err)
	}

}
