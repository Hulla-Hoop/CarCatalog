package endpoint

import "net/http"

type Endpoint interface {
	Insert(w http.ResponseWriter, r *http.Request)
}
