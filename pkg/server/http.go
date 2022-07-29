package server

import (
	"net/http"
	"time"

	"github.com/marlaone/website/pkg/config"
	"github.com/marlaone/website/pkg/contents"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type HttpServer struct {
	logger *zap.Logger
}

func NewHttpServer(logger *zap.Logger) *HttpServer {
	return &HttpServer{
		logger: logger,
	}
}

func (s *HttpServer) Serve() error {

	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(NewLoggerMiddleware(s.logger, &LoggerOpts{WithReferer: true, WithUserAgent: true}))
	r.Use(middleware.Recoverer)
	r.Use(middleware.Compress(5, "gzip"))
	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	httpFileServer := http.FileServer(http.Dir("./web/dist"))

	// @TODO refactor to a better place
	cacheControlHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "max-age=31536000")
		httpFileServer.ServeHTTP(w, r)
	}

	if config.IsDebug() {
		r.Mount("/debug", middleware.Profiler())
	}
	r.Handle("/public/*", http.StripPrefix("/public/", http.HandlerFunc(cacheControlHandler)))
	r.Handle("/_marla/*", http.StripPrefix("/_marla/", http.HandlerFunc(cacheControlHandler)))
	r.Handle("/*", contents.Handler(s.logger))

	return http.ListenAndServe(":"+viper.GetString(config.KeyHttpPort), r)
}
