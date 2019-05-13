package server

import (
	"encoding/json"
	"net/http"
)

// NotFound handles routes that do not exist
func (s *Srv) notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Oops! We haven't built anything here... yet ;)")
	return
}
