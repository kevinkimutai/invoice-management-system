package handlers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kevinkimutai/invoice-management-system/internal/adapters/queries"
	"github.com/kevinkimutai/invoice-management-system/internal/domain"
	"github.com/kevinkimutai/invoice-management-system/internal/ports"
)

type CompanyService struct {
	api ports.CompanyApiPort
}

func NewCompanyService(api ports.CompanyApiPort) *CompanyService {
	return &CompanyService{api: api}
}
func (s *CompanyService) GetAllCompanies(c *fiber.Ctx) error {
	//Get Query Params
	m := c.Queries()

	//Bind To CustomerParams
	params := domain.CheckParams(m)

	//Get All Customers API
	companies, err := s.api.GetAllCompanies(params)
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})

	}
	return c.Status(200).JSON(
		domain.CompaniesResponse{
			StatusCode:    200,
			Message:       "Successfully retrieved companies",
			Page:          companies.Page,
			NumberOfPages: companies.NumberOfPages,
			Total:         companies.Total,
			Data:          companies.Data,
		})
}

func (s *CompanyService) CreateCompany(c *fiber.Ctx) error {

	//Get Logo File
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	// Save file to public/images directory:
	destination := fmt.Sprintf("./public/images/%d-%s", time.Now().Unix(), file.Filename)
	if err := c.SaveFile(file, destination); err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	//Get UserID
	user := c.Locals("customer")
	fmt.Println(user)

	userTypeAssert, ok := user.(queries.User)
	if !ok {
		fmt.Println("Type assertion failed, user is not of type queries.User")

	}

	//Receive company details
	companyName := c.FormValue("company_name")
	if companyName == "" {
		return c.Status(400).JSON(
			domain.ErrorResponse{
				StatusCode: 400,
				Message:    "Missing company name",
			})
	}

	//Save To DB
	com := domain.Company{
		LogoUrl:     destination,
		UserID:      userTypeAssert.UserID,
		CompanyName: companyName,
	}

	savedCompany, err := s.api.CreateCompany(com)
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})

	}

	return c.Status(201).JSON(
		domain.CompanyResponse{
			StatusCode: 201,
			Message:    "Successfully created company",
			Data:       savedCompany})

}
