package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/defilippomattia/go-rest-template/internal/config"
	"github.com/defilippomattia/go-rest-template/internal/db"
	"github.com/defilippomattia/go-rest-template/internal/endpoints/healthz"
	"github.com/defilippomattia/go-rest-template/internal/logging"

	"github.com/go-chi/chi/v5"
)

func run(ctx context.Context, w io.Writer, args []string) error {
	var cfg config.Config
	if err := config.LoadConfig(w, "internal/config/config.toml", &cfg); err != nil {
		return err
	}

	if err := config.ValidateConfig(cfg); err != nil {
		return err
	}
	logger := logging.GetLogger(w, cfg.Logging.Level)

	logger.Debug().Msg("Logger initialized")
	logger.Debug().Msgf("Creating dbCredentials struct...")
	dbCredentials := db.Credentials{
		Type:     "postgres",
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		User:     cfg.DB.User,
		Password: cfg.DB.Password,
		Database: cfg.DB.Database,
	}
	logger.Debug().Msg("dbCredentials struct created")

	logger.Debug().Msgf("Connecting to the database...")
	dbconn, err := db.GetConnectionPool(dbCredentials)
	if err != nil {
		logger.Error().Err(err).Msg("Error connecting to the database, exiting server.")
		return err
	}
	defer dbconn.Close()
	logger.Debug().Msg("Connected to the database!")
	logger.Debug().Msg("Creating server dependencies...")
	sd := &config.ServerDeps{
		Db:     dbconn,
		Logger: logger,
	}

	rt := chi.NewRouter()
	healthz.RegisterRoutes(rt, sd)

	hostAndPort := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	err = http.ListenAndServe(hostAndPort, rt)
	if err != nil {
		logger.Error().Err(err).Msg("Error starting server")
		return err
	}
	return nil

}

func main() {
	logfile, err := logging.CreateOrOpenLogFile()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create or open log file: %s\n", err)
		os.Exit(1)
	}
	defer logfile.Close()

	ctx := context.Background()

	fmt.Fprintf(logfile, "####################################################### Starting server... ####################################################### \n")

	if err := run(ctx, logfile, os.Args); err != nil {
		fmt.Fprintf(logfile, "ERROR: %s\n", err)
		os.Exit(1)
	}
}
