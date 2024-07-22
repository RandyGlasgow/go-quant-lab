package http_utils

import "net/http"

func HttpStandardError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), http.StatusInternalServerError)
}
