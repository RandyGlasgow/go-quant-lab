package http_utils

import "net/http"

func HttpCustomError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusBadRequest)
}
