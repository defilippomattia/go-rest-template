package loglevel

import (
	"net/http"

	"github.com/defilippomattia/go-rest-template/internal/config"
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(rt chi.Router, sd *config.ServerDeps) {
	rt.Get("/api/v1/loglevel", changeLogLevel(sd))
}

func changeLogLevel(sd *config.ServerDeps) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//todo...
		levelStr := r.URL.Query().Get("level")
		sd.Logger.Info().Msgf("Trying to change log level to %s", levelStr)

	}
}
