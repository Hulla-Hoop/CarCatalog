package cataloge

import (
	"encoding/json"
	"net/http"
)

func (e *endpoint) Filter(w http.ResponseWriter, r *http.Request) {

	reqID, ok := r.Context().Value("reqID").(string)
	if !ok {
		reqID = ""
	}

	limit := r.URL.Query().Get("limit")
	offset := r.URL.Query().Get("offset")
	field := r.URL.Query().Get("field")
	value := r.URL.Query().Get("value")
	operation := r.URL.Query().Get("operator")

	cars, err := e.s.Filter(reqID, limit, offset, field, value, operation)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(cars)

}
