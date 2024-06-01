package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/kevinkimutai/invoice-management-system/internal/ports"
)

type ServerAdapter struct {
	port    string
	auth    ports.AuthHandlerPort
	user    ports.UserHandlerPort
	company ports.CompanyHandlerPort
	invoice ports.InvoiceHandlerPort
}

func New(
	port string,
	auth ports.AuthHandlerPort,
	user ports.UserHandlerPort,
	company ports.CompanyHandlerPort,
	invoice ports.InvoiceHandlerPort,
) *ServerAdapter {

	return &ServerAdapter{
		port:    port,
		auth:    auth,
		user:    user,
		company: company,
		invoice: invoice,
	}

}

func (s *ServerAdapter) Run() {
	//Initialize Fiber
	app := fiber.New()

	//Logger Middleware
	app.Use(logger.New())

	// Define routes
	app.Route("/api/v1/user", s.UserRouter)
	app.Route("/api/v1/company", s.CompanyRouter)
	app.Route("/api/v1/invoice", s.InvoiceRouter)

	app.Listen(":" + s.port)
}
