package httpserver

import "log/slog"

type HttpServer struct {
	playlistService playlistService
	songService songService
	log *slog.Logger
}

func NewHttpServer(log *slog.Logger, playlistService playlistService, songService songService) HttpServer {
	return HttpServer{
		playlistService: playlistService,
		songService: songService,
		log: log,
	}
}