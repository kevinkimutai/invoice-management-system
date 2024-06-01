package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kevinkimutai/invoice-management-system/internal/domain"
	"github.com/kevinkimutai/invoice-management-system/internal/ports"
)

type UserService struct {
	api ports.UserApiPort
}

func NewUserService(api ports.UserApiPort) *UserService {
	return &UserService{api: api}
}
func (s *UserService) GetAllUsers(c *fiber.Ctx) error {
	//Get Query Params
	m := c.Queries()

	//Bind To CustomerParams
	params := domain.CheckParams(m)

	//Get All Customers API
	users, err := s.api.GetAllUsers(params)
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})

	}
	return c.Status(200).JSON(
		domain.UsersResponse{
			StatusCode:    200,
			Message:       "Successfully retrieved users",
			Page:          users.Page,
			NumberOfPages: users.NumberOfPages,
			Total:         users.Total,
			Data:          users.Data,
		})

}
