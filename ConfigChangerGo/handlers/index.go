package handlers

import (
	"net/http"
)

// swagger:route GET / index
// Returns OK if there's no problem
// responses:
//	200: RatingResponse

// Index returns OK handles GET requests
func (p *APIContext) Index(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(200)
}
