package healthz

import (
	"net/http"

	"github.com/defilippomattia/go-rest-template/internal/config"
	"github.com/defilippomattia/go-rest-template/internal/reqresp"
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(rt chi.Router, sd *config.ServerDeps) {
	rt.Get("/api/v1/healthz", checkHealth(sd))
}

func checkHealth(sd *config.ServerDeps) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := reqresp.Encode(w, r, http.StatusOK, "OK")
		if err != nil {
			sd.Logger.Error().Err(err).Msg("Error encoding response")
		}
	}
}
