package server

import "github.com/gofiber/fiber/v2"

func (s *ServerAdapter) CompanyRouter(api fiber.Router) {
	api.Get("/", s.auth.IsAuthenticated, s.auth.AllowedRoles("Admin"), s.company.GetAllCompanies)
	api.Post("/", s.auth.IsAuthenticated, s.company.CreateCompany)
}
