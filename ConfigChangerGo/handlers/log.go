package handlers

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// swagger:route GET /log Log Log
// Return 200 and logs something at all levels
// responses:
//	200: OK
//	404: errorResponse

// Log handles GET requests
func (ctx *APIContext) Log(rw http.ResponseWriter, r *http.Request) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	fmt.Println("-----------------------------------------")
	log.Debug().Msg("Debug Log")
	log.Info().Msg("Info Log")
	log.Warn().Msg("Warn Log")
	log.Error().Msg("Error Log")
	fmt.Println("Fatal logging causes the program to quit, that's why we omitted fatal logging")
	rw.WriteHeader(http.StatusOK)
}
