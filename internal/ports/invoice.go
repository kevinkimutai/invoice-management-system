package ports

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kevinkimutai/invoice-management-system/internal/adapters/queries"
	"github.com/kevinkimutai/invoice-management-system/internal/domain"
	"github.com/kevinkimutai/invoice-management-system/internal/utils"
)

type InvoiceRepoPort interface {
	CreateInvoice(utils.InvoiceParams) (domain.Invoice, error)
	GetInvoice(string) (domain.Invoice, error)
	GetInvoiceData(string) (queries.GetInvoiceDataByInvoiceIDRow, error)
	GetInvoiceItemsData(string) ([]queries.GetInvoiceItemDataByInvoiceIDRow, error)
}

type InvoiceHandlerPort interface {
	CreateInvoice(c *fiber.Ctx) error
	GetInvoiceByID(c *fiber.Ctx) error
	DownloadInvoice(c *fiber.Ctx) error
}

type InvoiceApiPort interface {
	CreateInvoice(domain.Invoice) (domain.Invoice, error)
	GetInvoiceByID(string) (domain.Invoice, error)
	DownloadInvoice(string) error
}
