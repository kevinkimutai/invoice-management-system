package ports

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kevinkimutai/invoice-management-system/internal/domain"
	"github.com/kevinkimutai/invoice-management-system/internal/utils"
)

type CompanyRepoPort interface {
	GetCompanies(params utils.LimitParams) (domain.CompaniesFetch, error)
	CreateCompany(company domain.Company) (domain.Company, error)
}

type CompanyHandlerPort interface {
	GetAllCompanies(c *fiber.Ctx) error
	CreateCompany(c *fiber.Ctx) error
}

type CompanyApiPort interface {
	GetAllCompanies(params domain.Params) (domain.CompaniesFetch, error)
	CreateCompany(company domain.Company) (domain.Company, error)
}
