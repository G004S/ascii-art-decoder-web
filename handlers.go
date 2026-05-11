package main
import ("net/http"
)

type PageData struct{
	Input string
	Result string
}

//Handler for homepage
func (app *application) HomeHandler(w http.ResponseWriter, r *http.Request) {
	//executes HTML template (Writer, HTMLname, DATA)
	app.templates.ExecuteTemplate(w, "home.html", nil)
	
}

//Handler for decoder
func (app *application) DecoderHandler(w http.ResponseWriter, r *http.Request) {
	//validate request
	if r.Method == http.MethodPost{
		r.ParseForm()
		input:= r.Form.Get("input")
		data:=PageData{}

		//incorrect input -> error 400
		if input == ""{
			w.WriteHeader(http.StatusBadRequest)
			app.templates.ExecuteTemplate(w, "decoder.html", nil)
			return
		}
		res, err := Decode(input)
		//incorrect input -> error 400
		if err != nil{
			w.WriteHeader(http.StatusBadRequest)
			app.templates.ExecuteTemplate(w, "decoder.html", nil)
			return
		}
		data.Result = res
		data.Input = input
		//correct input -> 202 accepted
		w.WriteHeader(http.StatusAccepted)
		app.templates.ExecuteTemplate(w, "decoder.html", data)
	}else{
		app.templates.ExecuteTemplate(w, "decoder.html", nil)
	}

	
	
}

//Handler for encoder
func (app *application) EncoderHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost{
		r.ParseForm()
		input:= r.Form.Get("input")
		data:=PageData{}

		if input==""{
			w.WriteHeader(http.StatusBadRequest)
			app.templates.ExecuteTemplate(w,"encoder.html",nil)
			return
		}
		res,err:=Encode(input)
		if err!=nil{
			w.WriteHeader(http.StatusBadRequest)
			app.templates.ExecuteTemplate(w,"encoder.html",nil)
		}
		data.Result=res
		data.Input=input
		w.WriteHeader(http.StatusAccepted)
			app.templates.ExecuteTemplate(w,"encoder.html",data)
	}else{
			app.templates.ExecuteTemplate(w, "encoder.html", nil)
	}
	
}
