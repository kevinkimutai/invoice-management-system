package server

import "github.com/gofiber/fiber/v2"

func (s *ServerAdapter) UserRouter(api fiber.Router) {
	api.Get("/", s.auth.IsAuthenticated, s.auth.AllowedRoles("Admin"), s.user.GetAllUsers)
}
