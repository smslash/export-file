package pkg

import (
	"html/template"
	"log"
	"net/http"
)

func NotFound(w http.ResponseWriter) {
	tmp, err := template.ParseFiles("./ui/html/error404.html")
	if err != nil {
		log.Println(err.Error())
		InternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusNotFound)

	err = tmp.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		InternalServerError(w)
		return
	}
}

func MethodNotAllowed(w http.ResponseWriter) {
	tmp, err := template.ParseFiles("./ui/html/error405.html")
	if err != nil {
		log.Println(err.Error())
		InternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)

	err = tmp.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		InternalServerError(w)
		return
	}
}

func BadRequest(w http.ResponseWriter) {
	tmp, err := template.ParseFiles("./ui/html/error400.html")
	if err != nil {
		log.Println(err.Error())
		InternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusBadRequest)

	err = tmp.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		InternalServerError(w)
		return
	}
}

func InternalServerError(w http.ResponseWriter) {
	tmp, err := template.ParseFiles("./ui/html/error500.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusInternalServerError)

	err = tmp.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
