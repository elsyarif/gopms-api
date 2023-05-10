package http

import (
	"context"
	"fmt"
	"github.com/elsyarif/pms-api/internal/infrastructure"
	"github.com/elsyarif/pms-api/pkg/helper/log"
	"github.com/elsyarif/pms-api/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type AppService struct {
}

type Server struct {
	engine          *gin.Engine
	httpAddr        string
	shutdownTimeout time.Duration
	db              *sqlx.DB
}

func serverContext(ctx context.Context) context.Context {
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		<-c
		cancel()
	}()

	return ctx
}

func NewServer(
	ctx context.Context,
	host string,
	port uint,
	shutdownTimeout time.Duration,
	db *sqlx.DB) (context.Context, Server) {
	svr := Server{
		engine:          gin.Default(),
		httpAddr:        fmt.Sprintf("%s:%d", host, port),
		shutdownTimeout: shutdownTimeout,
		db:              db,
	}

	// gin.SetMode(gin.ReleaseMode)
	svr.engine.Use(middleware.CorsMiddleware())
	svr.engine.Use(middleware.ErrorHandler)
	infrastructure.Container(db, svr.engine)
	svr.engine.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "404 page not found",
		})
	})
	return serverContext(ctx), svr
}

func (s *Server) Run(ctx context.Context) error {
	log.Info("Server running", logrus.Fields{"server": s.httpAddr})

	svr := http.Server{
		Addr:    s.httpAddr,
		Handler: s.engine,
	}

	go func() {
		if err := svr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Failed to start server", logrus.Fields{"error": err.Error()})
		}
	}()

	<-ctx.Done()
	log.Info("Received shutdown signal", logrus.Fields{})
	ctxShutdown, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	// Shutdown server
	return svr.Shutdown(ctxShutdown)
}
