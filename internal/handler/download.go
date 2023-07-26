package handler

import (
	"net/http"
	"strconv"

	"git/ssengerb/ascii-art-web/pkg"
)

func Download(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		pkg.MethodNotAllowed(w)
		return
	}

	data := r.FormValue("data")

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.Header().Set("Content-Disposition", "attachment; filename=output.txt")
	_, err := w.Write([]byte(data))
	if err != nil {
		pkg.InternalServerError(w)
		return
	}
}
