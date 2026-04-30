package main

import (
    /* "fmt" */
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
	//validate request
	if r.Method == http.MethodPost{
		r.ParseForm()
		input:= r.Form.Get("input")
		data:=PageData{}

		//incorrect input -> error 400
		if input == ""{
			w.WriteHeader(http.StatusBadRequest)
			tmpl.ExecuteTemplate(w, "decoder.html", nil)
			return
		}
		res, err := Decode(input)
		//incorrect input -> error 400
		if err != nil{
			w.WriteHeader(http.StatusBadRequest)
			tmpl.ExecuteTemplate(w, "decoder.html", nil)
			return
		}
		data.Result = res
		data.Input = input
		//correct input -> 202 accepted
		w.WriteHeader(http.StatusAccepted)
		tmpl.ExecuteTemplate(w, "decoder.html", data)
	}else{
		//incorrect request -> error 400
		w.WriteHeader(http.StatusBadRequest)
		tmpl.ExecuteTemplate(w, "decoder.html", nil)
	}

	
	
}

//Handle for encoder
func EncoderHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost{
		r.ParseForm()
		input:= r.Form.Get("input")
		data:=PageData{}

		if input==""{
			w.WriteHeader(http.StatusBadRequest)
			tmpl.ExecuteTemplate(w,"encoder.html",nil)
			return
		}
		res,err:=Encode(input)
		if err!=nil{
			w.WriteHeader(http.StatusBadRequest)
			tmpl.ExecuteTemplate(w,"encoder.html",nil)
		}
		data.Result=res
		data.Input=input
		w.WriteHeader(http.StatusAccepted)
		tmpl.ExecuteTemplate(w,"encoder.html",data)
	}else{
		w.WriteHeader(http.StatusBadRequest)
		tmpl.ExecuteTemplate(w, "decoder.html", nil)
	}
	
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
