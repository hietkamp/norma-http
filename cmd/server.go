package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hietkamp/norma-http/internal/handlers"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Server struct {
	HttpServer *http.Server
}

func init() {
	// Force log's color
	gin.ForceConsoleColor()
	gin.SetMode(gin.ReleaseMode)
	if gin.IsDebugging() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func useLoggingMiddleware(r *gin.Engine) {
	r.Use(func(c *gin.Context) {
		// Read from body and write here again.
		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
			log.Debug().Msgf("Header: %s", c.Request.Header)
			log.Debug().Msgf("Body: %s", string(bodyBytes))
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		c.Next()
	})
}
func NewServer() *Server {
	router := gin.Default()
	// The order of loading the middleware matters, cors need to be first
	router.Use(cors.Default())
	useLoggingMiddleware(router)

	// Ping test
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	router.POST("/messaging", handlers.HandleMessage)
	router.POST("/accesstoken", handlers.HandleAccessToken)
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	server := Server{
		HttpServer: &http.Server{
			Addr: "0.0.0.0:8080",
			// Good practice to set timeouts to avoid Slowloris attacks.
			WriteTimeout: time.Second * 60,
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Second * 60,
			Handler:      router,
		},
	}
	return &server
}

func (s *Server) Serve() {
	// Initializing the httpserver in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		log.Info().Msgf("HTTP Server Listening on %s", s.HttpServer.Addr)
		if err := s.HttpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Msgf("Error while listening: %s", err)
		}
	}()
}
