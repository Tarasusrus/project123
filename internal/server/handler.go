package server

import (
	"BaseApi/internal/logger"
	"BaseApi/internal/server/middleware"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

// AppCfg - конфигурационные данные приложения
type AppCfg struct {
	Mode      string `mapstructure:"APP_MODE" default:"release"` // может быть "dev" и "release". "dev" открывает токен обхода
	Host      string `mapstructure:"APP_HOST" default:"localhost"`
	Port      string `mapstructure:"APP_PORT" default:"8085"`
	ApiPrefix string `mapstructure:"API_PREFIX" default:""`
}

// Handler - структура хендлера
type Handler struct {
	Router *mux.Router
	Server *http.Server
	logger logger.Logger
}

// NewHandler - возвращает хендлер
func NewHandler(cfg *AppCfg, logger logger.Logger) *Handler {
	h := &Handler{
		logger: logger,
	}
	h.Router = mux.NewRouter()
	h.mapRoutes(cfg.ApiPrefix)

	addr := cfg.Host + ":" + cfg.Port
	logger.Info("app_start", "addr", addr)

	h.Server = &http.Server{
		Addr:    addr,
		Handler: h.Router,
	}

	return h
}

func (h *Handler) mapRoutes(prefix string) {
	prefixRouter := h.Router.PathPrefix(prefix).Subrouter()

	appRouter := prefixRouter.PathPrefix("/api/v1").Subrouter()

	// @Summary Health Check
	// @Description Проверка доступности API
	// @ID health-check
	// @Produce plain
	// @Success 200 {string} string "OK"
	// @Router /api/v1/ [head]
	appRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods(http.MethodHead)

	appRouter.Use(middleware.RequestID)
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
