package httpserver

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/serj213/music-playlist/internal/app/common/server"
	"github.com/serj213/music-playlist/internal/app/common/sl"
)

func (p HttpServer) AddSong(w http.ResponseWriter, r *http.Request) {

	log := p.log.With(slog.String("methor", "addSong"))

	var songReq AddSongRequest
	queryParams := r.URL.Query()

	playlistIdParam := queryParams.Get("playlist_id")

	if playlistIdParam == "" {
		server.BadRequest("invalid playlist_id", w, r)
		return
	}

	playlistId, err := strconv.Atoi(playlistIdParam)
	if err != nil {
		server.BadRequest("invalid playlist_id", w, r)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&songReq); err != nil {
		fmt.Println("failed decode add song request: %w", err )
		server.BadRequest("invalid request", w, r)
		return
	}

	defer r.Body.Close()

	if err := songReq.Validate(); err != nil {
		server.BadRequest(fmt.Sprintf("invalid request: %v", err), w, r)
		return
	}

	songDomain, err := songToDomain(songReq)
	if err != nil {
		log.Info("failed convert song to domain: %w", sl.Error(err))
		server.InternalServer("server error", w, r)
		return
	}

	songId, err := p.songService.AddSong(r.Context(), songDomain, playlistId)
	if err != nil {
		log.Info("failed add song: %w", sl.Error(err))
		server.InternalServer("server error", w, r)
		return
	}

	server.ResponseOk(AddSongResponse{
		Id: songId,
	}, w, r)

	log.Info("success add song")
}