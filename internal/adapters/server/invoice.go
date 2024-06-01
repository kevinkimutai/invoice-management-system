package server

import "github.com/gofiber/fiber/v2"

func (s *ServerAdapter) InvoiceRouter(api fiber.Router) {
	api.Post("/", s.auth.IsAuthenticated, s.invoice.CreateInvoice)
	api.Get("/:invoiceID", s.auth.IsAuthenticated, s.invoice.GetInvoiceByID)
	api.Get("/:invoiceID/download", s.auth.IsAuthenticated, s.invoice.DownloadInvoice)

}
