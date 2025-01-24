package httpserver

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/serj213/music-playlist/internal/app/common/server"
	"github.com/serj213/music-playlist/internal/app/common/sl"
)


func (p HttpServer) AddSong (log *slog.Logger, w http.ResponseWriter, r *http.Request) {
	var songRequest SongRequest

	log = log.With(slog.String("handler", "add song"))

	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&songRequest); err != nil {
		log.Error("failed decode body", sl.Error(err.Error()))
		server.BadRequest("invalid request", err, w, r)
		return
	}

	if err := songRequest.Validate(); err != nil {
		log.Error("error validate", sl.Error(err.Error()))
		return
	}
}