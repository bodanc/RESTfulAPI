package handlers

import "net/http"

// RootHandler handles the root route
func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Path not found"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Running APIv1\n"))
}
