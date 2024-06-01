package ports

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kevinkimutai/invoice-management-system/internal/adapters/queries"
	"github.com/kevinkimutai/invoice-management-system/internal/domain"
	"github.com/kevinkimutai/invoice-management-system/internal/utils"
)

type UserRepoPort interface {
	GetUsers(utils.LimitParams) (domain.UsersFetch, error)
}

type DBRepoPort interface {
	CreateUser(queries.CreateUserParams) (queries.User, error)
	GetUserByEmail(string) (queries.User, error)
}

type UserHandlerPort interface {
	GetAllUsers(c *fiber.Ctx) error
}

type UserApiPort interface {
	GetAllUsers(domain.Params) (domain.UsersFetch, error)
}
