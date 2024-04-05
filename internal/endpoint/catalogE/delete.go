package cataloge

import (
	"encoding/json"
	"net/http"
)

func (e *endpoint) Delete(w http.ResponseWriter, r *http.Request) {
	reqID, ok := r.Context().Value("reqID").(string)
	if !ok {
		reqID = ""
	}
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "некорректные данные", http.StatusBadRequest)
		return
	}
	car, err := e.s.Delete(reqID, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(car)
}
