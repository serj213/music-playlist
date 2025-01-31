package httpserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/serj213/music-playlist/internal/app/common/server"
)


func (p HttpServer) CreatePlaylist(w http.ResponseWriter, r *http.Request) {
	var playlistReq PlaylistRequest

	if err := json.NewDecoder(r.Body).Decode(&playlistReq); err != nil {
		server.BadRequest("invalid request", w, r)
		return
	}

	defer r.Body.Close()

	playlistDomain, err := playlistToDomain(playlistReq)
	if err != nil {
		server.InternalServer("server error", w, r)
		return
	}

	playlist, err := p.playlistService.CreatePlaylist(r.Context(), playlistDomain)

	if err != nil {
		fmt.Printf("failed create playlist: %w", err)
		server.InternalServer("server error", w, r)
		return
	}

	server.ResponseOk(PlaylistResponse{
		Status: "Ok",
		Playlist: playlistRes{
			Id: playlist.ID(),
			Title: playlist.Title(),
		},
	}, w, r)

}


/*

func (p HttpServer) AddSong (w http.ResponseWriter, r *http.Request) {
	var songRequest SongRequest
	if err := json.NewDecoder(r.Body).Decode(&songRequest); err != nil {
		server.BadRequest("invalid request", err, w, r)
		return
	}

	defer r.Body.Close()

	if err := songRequest.Validate(); err != nil {
		return
	}

	songDomain := domain.NewSong(domain.NewSongData{
		Title: "Yagofffdaf",
		Artist: "авав",
		Duration: 2113131313,
	})

	_, err := p.playlistService.AddSong(r.Context(), songDomain)
	if err != nil {
		fmt.Println("ывыфвфв: %w", err)
	}

}

*/