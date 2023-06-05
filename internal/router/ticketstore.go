package router

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/RemyRanger/go-etl-csv/internal/common/config"
	"github.com/RemyRanger/go-etl-csv/internal/common/logs"
	"github.com/RemyRanger/go-etl-csv/internal/common/server/middleware"
	"github.com/RemyRanger/go-etl-csv/internal/services/tickets"
	ticketsPorts "github.com/RemyRanger/go-etl-csv/internal/services/tickets/handler"
	ticketsRepo "github.com/RemyRanger/go-etl-csv/internal/services/tickets/repository"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"gorm.io/gorm"
)

// TicketstoreRouter : TicketstoreRouter
type TicketstoreRouter struct {
	db             *gorm.DB
	httpServer     *http.Server
	ticketsService tickets.Service
}

// NewTicketstoreRouter : NewTicketstoreRouter
func NewTicketstoreRouter(db *gorm.DB) *TicketstoreRouter {
	return &TicketstoreRouter{
		db:             db,
		ticketsService: tickets.NewService(ticketsRepo.NewRepository(db)),
	}
}

// Serve : run service
func (app *TicketstoreRouter) Serve() {

	r := chi.NewRouter()
	const path = "/v1"

	setMiddlewares(r, path, app.db)

	ticketsPorts.HandlerWithOptions(ticketsPorts.NewHTTPHandler(app.ticketsService), ticketsPorts.ChiServerOptions{BaseURL: path, BaseRouter: r})

	app.httpServer = &http.Server{
		Handler:           r,
		Addr:              config.Conf.Srv.Addr,
		ReadHeaderTimeout: 100 * time.Millisecond,
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Info().Msg("API shutting down")

			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			cancel()

			// start http shutdown
			if err := app.httpServer.Shutdown(ctx); err != nil {
				log.Fatal().Err(err).Msg("Error shutting down HTTPS server")
			}

			// verify, in worst case call cancel via defer
			select {
			case <-time.After(31 * time.Second):
				log.Info().Msg("Not all connections done")
			case <-ctx.Done():

			}
		}
	}()
	err := app.httpServer.ListenAndServeTLS(config.Conf.Srv.Certificate, config.Conf.Srv.Key)
	if err != nil && err != http.ErrServerClosed {
		log.Error().Err(err).Msg("Error starting HTTPS server")
	}
}

func setMiddlewares(router *chi.Mux, path string, db *gorm.DB) {

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"POST"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	})

	requestLogsMiddleware := middleware.NewRequestLogsMiddleware(db)

	router.Use(
		// middlewares
		corsMiddleware.Handler,

		// healthcheck
		chiMiddleware.Heartbeat(path+"/healthcheck"),

		// custom middlewares
		requestLogsMiddleware.Handler,

		// custom chi logger
		logs.Logger,
	)

}
