package httpserver

type HttpServer struct {
	playlistService playlistService
}

func NewHttpServer(playlistService playlistService) HttpServer {
	return HttpServer{
		playlistService: playlistService,
	}
}