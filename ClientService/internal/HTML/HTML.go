package gui_html

import (
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	URL = "/"
)

type Handler struct {
}

func (h *Handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, URL, h.MainPage)
}

//TODO: Set HTML template to config file?
func (h *Handler) MainPage(w http.ResponseWriter, _ *http.Request) {
	ts, err := template.ParseFiles("./internal/HTML/main.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
