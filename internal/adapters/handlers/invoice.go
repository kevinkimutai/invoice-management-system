package handlers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kevinkimutai/invoice-management-system/internal/adapters/queries"
	"github.com/kevinkimutai/invoice-management-system/internal/domain"
	"github.com/kevinkimutai/invoice-management-system/internal/ports"
	"github.com/kevinkimutai/invoice-management-system/internal/utils"
)

type InvoiceService struct {
	api ports.InvoiceApiPort
}

func NewInvoiceService(api ports.InvoiceApiPort) *InvoiceService {
	return &InvoiceService{
		api: api,
	}
}

func (s *InvoiceService) CreateInvoice(c *fiber.Ctx) error {
	invoice := domain.Invoice{}

	//Bind To struct
	if err := c.BodyParser(&invoice); err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	//Check Missing Values
	in, err := domain.NewInvoice(invoice)
	if err != nil {
		return c.Status(400).JSON(
			domain.ErrorResponse{
				StatusCode: 400,
				Message:    err.Error(),
			})
	}

	//Generate UUID
	invoiceId := time.Now().Unix()
	invoiceIdStr := utils.ConvertInt64ToString(invoiceId)

	//Get UserID
	user := c.Locals("customer")
	//fmt.Println(user)

	userTypeAssert, ok := user.(queries.User)
	if !ok {
		fmt.Println("Type assertion failed, user is not of type queries.User")

	}

	newInvoice := domain.Invoice{
		ID:             in.ID,
		InvoiceID:      invoiceIdStr,
		UserID:         userTypeAssert.UserID,
		CompanyID:      in.CompanyID,
		Address:        in.Address,
		AccountNumber:  in.AccountNumber,
		TotalAmount:    in.TotalAmount,
		InvoiceDate:    in.InvoiceDate,
		InvoiceDueDate: in.InvoiceDueDate,
		InvoiceType:    in.InvoiceType,
		Items:          in.Items,
	}

	//Save To DB
	invoice, err = s.api.CreateInvoice(newInvoice)
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	return c.Status(201).JSON(
		domain.InvoiceResponse{
			StatusCode: 201,
			Message:    "Successfully created invoice",
			Data:       invoice})
}

func (s *InvoiceService) GetInvoiceByID(c *fiber.Ctx) error {
	invoiceID := c.Params("invoiceID")

	invoice, err := s.api.GetInvoiceByID(invoiceID)
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	return c.Status(200).JSON(
		domain.InvoiceResponse{
			StatusCode: 200,
			Message:    "Success",
			Data:       invoice})
}

func (s *InvoiceService) DownloadInvoice(c *fiber.Ctx) error {
	invoiceID := c.Params("invoiceID")

	err := s.api.DownloadInvoice(invoiceID)

	return err
}
