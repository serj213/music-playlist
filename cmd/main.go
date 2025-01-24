package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/serj213/music-playlist/internal/app/config"
	pgrepo "github.com/serj213/music-playlist/internal/app/repository/pgRepo"
	"github.com/serj213/music-playlist/internal/app/service"
	httpserver "github.com/serj213/music-playlist/internal/app/transport/http_server"
	"github.com/serj213/music-playlist/internal/pkg/pg"
)

const (
	local = "local"
	dev = "develop"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}


func run()  error {
	cfg, err := config.Read()

	if err != nil {
		return err
	}

	log := setupLogger(cfg.Env)
	log = log.With(slog.String("env", cfg.Env))

	log.Info("logger enabled")

	pgDb, err  := pg.Deal(cfg.DSN)

	if err != nil {
		return err
	}

	log.Info("success connect database")

	playlistRepo := pgrepo.NewPgRepo(pgDb)

	playlistService := service.NewPlaylistService(playlistRepo)
	
	httpServer := httpserver.NewHttpServer(playlistService)

	rounter := mux.NewRouter()

	log.Info("init router")
	rounter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("playlist API"))
	})
	
	_ = httpServer

	return nil
}

func setupLogger(env string) *slog.Logger {

	var log *slog.Logger

	switch env {
	case local:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case dev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	default:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}