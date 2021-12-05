package cmd

import (
	"fmt"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pedromsmoreira/app-proxy/proxies"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

// Server struct, should have the needed dependencies
type Server struct {
	app *fiber.App
	wg  sync.WaitGroup
}

// NewServer method to init a new server instance
func NewServer() *Server {
	app := fiber.New(fiber.Config{
		ReadTimeout: 5 * time.Second, // to make sure keep alive connections are closed on Shutdown()
	})

	app.Use(cors.New())
	app.Get("/dashboard", monitor.New())
	app.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("ping"))
	})

	return &Server{
		app: app,
		wg:  sync.WaitGroup{},
	}
}

// Start method to start the server configurations
func (s *Server) Start(address string, port string, prxies *proxies.Proxies) error {
	var err error
	api := s.app.Group("/api")
	proxies.AddRoutesTo(api, prxies)

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		fmt.Printf("API Server listening on %v:%v", address, port)
		err := s.app.Listen(fmt.Sprintf("%v:%v", address, port))
		if err != nil {
			fmt.Printf("API Server stopped listening due to %v", err)
		} else {
			fmt.Printf("API Server stopped listening")
		}
	}()

	return err
}

func (s *Server) Shutdown() error {
	err := s.app.Shutdown()

	if err == nil {
		s.wg.Wait()
	}

	return err
}
