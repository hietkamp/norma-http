package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
	"time"

	"github.com/hietkamp/norma-http/cmd"
	"github.com/rs/zerolog/log"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

const (
	Version = "v1.0.1"
	Banner  = `

╔═╗ ╔╗                
║║╚╗║║                
║╔╗╚╝║╔══╗╔═╗╔╗╔╗╔══╗ 
║║╚╗║║║╔╗║║╔╝║╚╝║╚ ╗║ 
║║ ║║║║╚╝║║║ ║║║║║╚╝╚╗
╚╝ ╚═╝╚══╝╚╝ ╚╩╩╝╚═══╝ %s
Norma Simple HTTP Server, %s

___________________________________________________________/\__/\__0>______

`
)

func main() {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	fmt.Printf(Banner, string(colorRed)+Version+string(colorReset), string(colorCyan)+"https://infosupport.com"+string(colorReset))
	fmt.Print("\n")

	srv := cmd.NewServer()
	srv.Serve()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Info().Msgf("Shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the httpserver it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.HttpServer.Shutdown(ctx); err != nil {
		log.Fatal().Msgf("Server forced to shutdown: %s\n", err)
	}
	log.Info().Msgf("Server exiting\n")
}
