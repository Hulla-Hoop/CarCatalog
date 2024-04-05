package cataloge

import (
	"carcatalog/internal/model"
	"encoding/json"
	"io"
	"net/http"
)

func (e *endpoint) Update(w http.ResponseWriter, r *http.Request) {
	reqID, ok := r.Context().Value("reqID").(string)
	if !ok {
		reqID = ""
	}
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "некорректные данные", http.StatusBadRequest)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var car model.Car

	err = json.Unmarshal(body, &car)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cars, err := e.s.Update(reqID, car, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)

}
