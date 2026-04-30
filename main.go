package main

import (
    "fmt"
    "net/http"
	"html/template"
	/* "os" */
)


type PageData struct{
	Input string
	Result string
	Zdarova string
}

//STEP 3: Load Templates
//template.Must is the same thing as if err!= nil. SO JUST AN ERROR HANDLER
var tmpl = template.Must(template.ParseGlob("./templates/*.html")) //ParseGLob loads templates IN BULK


//STEP 2: Make handles for each page
//Handle for homepage
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	//executes HTML template (Writer, HTMLname, DATA)
	tmpl.ExecuteTemplate(w, "index.html", nil)
	
}

//Handle for decoder
func DecoderHandler(w http.ResponseWriter, r *http.Request) {
	//request parser. Prepares request data.
	r.ParseForm()
	//get "input" field from parsed request
	input:= r.Form.Get("input")

	data:=PageData{}
	if input != ""{
		res, err := Decode(input)
		if err != nil{
			fmt.Println(err)
		}

		data.Result = res
		data.Input = input
	}
	tmpl.ExecuteTemplate(w, "decoder.html", data)
}

//Handle for encoder
func EncoderHandler(w http.ResponseWriter, r *http.Request) {
	data:= PageData{}
	r.ParseForm()
	input:= r.Form.Get("input")
	if input != ""{
		res,err:= Encode(input)
		if err != nil{
			fmt.Println(err)
		}
		data.Result = res
		data.Input = input
	}
	
	tmpl.ExecuteTemplate(w, "encoder.html", data)
}




func main() {
	//STEP 1: allocate the MUX (multiplexer). Sort of a REQEST distributor
	mux:= http.NewServeMux()

	//if mux recieves "/" it pulls a specific handle.
    mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/decoder", DecoderHandler)
	mux.HandleFunc("/encoder", EncoderHandler)

	//STEP 1.1: Create a server (:port, mux)
    http.ListenAndServe(":8080", mux)
}
