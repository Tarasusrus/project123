package server

import (
	"BaseApi/internal/server/middleware"
	"BaseApi/internal/server/middleware/logger"
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

// AppCfg - конфигурационные данные приложения
type AppCfg struct {
	Mode                  string `mapstructure:"APP_MODE" default:"release"` // может быть "dev" и "release". "dev" открывает токен обхода
	Host                  string `mapstructure:"APP_HOST" default:"localhost"`
	Port                  string `mapstructure:"APP_PORT" default:"8085"`
	ExternalURL           string `mapstructure:"APP_EXTERNAL_URL"`
	AccessTokenTtlMinutes int    `mapstructure:"APP_ACCESS_TOKEN_TTL_MINUTES" default:"600"`
	JwtSecret             string `mapstructure:"APP_JWT_SECRET"`
	ApiKey                string `mapstructure:"SERVICE_API_KEY"`
	ApiPrefix             string `mapstructure:"API_PREFIX" default:""`
}

// Handler - структура хендлера
type Handler struct {
	Router *mux.Router
	Server *http.Server
	logger *slog.Logger
}

// NewHandler - возвращает хендлер
func NewHandler(cfg *AppCfg, log *slog.Logger) *Handler {
	h := &Handler{
		logger: log,
	}
	h.Router = mux.NewRouter()
	h.mapRoutes(cfg.ApiPrefix)

	addr := cfg.Host + ":" + cfg.Port
	log.Info("Start on: ", addr)

	h.Server = &http.Server{
		Addr:    addr,
		Handler: h.Router,
	}

	return h
}

func (h *Handler) mapRoutes(prefix string) {
	prefixRouter := h.Router.PathPrefix(prefix).Subrouter()

	appRouter := prefixRouter.PathPrefix("/api/v1").Subrouter()

	appRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods(http.MethodHead)

	appRouter.Use(middleware.RequestID)
	appRouter.Use(logger.New(h.logger))
}

// Serve - запуск сервера и graceful shutdown
func (h *Handler) Serve() error {

	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println(err.Error())
		}
	}()

	interruptChannel := make(chan os.Signal, 1)
	signal.Notify(interruptChannel, os.Interrupt)
	<-interruptChannel
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	err := h.Server.Shutdown(ctx)
	if err != nil {
		return err
	}

	log.Println("shut down gracefully")
	return nil
}
