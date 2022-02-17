package main

import (
	"github.com/billhcmus/conduit/api/test"
	"github.com/billhcmus/conduit/api/v1/user"
	"github.com/billhcmus/conduit/config"
	"github.com/billhcmus/conduit/logger"
	"github.com/billhcmus/conduit/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	conf := config.ServerConfig{Addr: ":8086"}
	s := server.New(conf, server.Option1(1))

	v1 := s.Group("api/v1")
	v2 := s.Group("api/v2")

	user.RegisterUserRoutes(v1)
	test.RegisterTestRoutes(v2)

	go s.Start()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	s.Stop()
	if err := logger.GetInstance().Sync(); err != nil {
		log.Fatalf("Failed to flush log %v", err)
	}
}
