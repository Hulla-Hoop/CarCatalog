package cataloge

import (
	"carcatalog/internal/model"
	"encoding/json"
	"io"
	"net/http"
)

func (e *endpoint) Insert(w http.ResponseWriter, r *http.Request) {
	reqID, ok := r.Context().Value("reqID").(string)
	if !ok {
		reqID = ""
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var reg model.RegNums

	err = json.Unmarshal(body, &reg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cars, err := e.s.Insert(reqID, reg.RegNums)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(cars)

}
