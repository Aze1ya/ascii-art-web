package handlers

import (
	"html/template"
	"log"
	"net/http"

	"01.alem.school/git/Azel/ascii-art-web-dockerize/ascii-art/utils"
)

type ViewData struct {
	Out  string
	Desc string
}
type ErrorData struct {
	Errortxt    string
	Errorstatus int
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	if status := checkErrHome(r); status != 0 {
		errorPage(w, http.StatusText(status), status)
		log.Print(http.StatusText(status))
		return
	}

	tmpl, err := template.ParseFiles("ui/html/site.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Print(http.StatusText(http.StatusInternalServerError))
		return
	}
	tmpl.Execute(w, nil)
}

func AsciiPage(w http.ResponseWriter, r *http.Request) {
	if status := checkErrAsciiPage(r); status != 0 {
		errorPage(w, http.StatusText(status), status)
		log.Print(http.StatusText(status))
		return
	}

	name := r.FormValue("style")
	text := r.FormValue("message")
	newData := new(ViewData)
	var errstatus int

	if newData.Out, errstatus = utils.AsciiConverter(w, text, name); errstatus != 0 {
		errorPage(w, http.StatusText(errstatus), errstatus)
		log.Print(http.StatusText(errstatus))
		return
	}

	newData.Desc = "Your text was converted to ascii-art"
	tmpl, err := template.ParseFiles("ui/html/site.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Print(http.StatusText(http.StatusInternalServerError))
		return
	}
	tmpl.Execute(w, newData)
}

func errorPage(w http.ResponseWriter, Errortxt string, Errorstatus int) {
	newErrorData := new(ErrorData)
	newErrorData.Errortxt = Errortxt
	newErrorData.Errorstatus = Errorstatus

	w.WriteHeader(Errorstatus)

	tmpl, err := template.ParseFiles("ui/html/error.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Print(http.StatusText(http.StatusInternalServerError))
		return
	}
	tmpl.Execute(w, newErrorData)
}

func checkErrHome(r *http.Request) int {
	if r.URL.Path != "/" {
		return http.StatusNotFound
	}
	if r.Method != "GET" && r.Method != "OPTIONS" && r.Method != "HEAD" {
		return http.StatusMethodNotAllowed
	}
	return 0
}

func checkErrAsciiPage(r *http.Request) int {
	if r.Method != "POST" && r.Method != "OPTIONS" && r.Method != "HEAD" {
		return http.StatusMethodNotAllowed
	}

	name := r.FormValue("style")
	text := r.FormValue("message")

	if text == "" {
		return http.StatusBadRequest
	}
	if name != "standard" && name != "thinkertoy" && name != "shadow" {
		return http.StatusBadRequest
	}
	return 0
}
