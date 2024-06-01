package pdf

import (
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jung-kurt/gofpdf"
	"github.com/kevinkimutai/invoice-management-system/internal/adapters/queries"
	"github.com/kevinkimutai/invoice-management-system/internal/utils"
)

type PDFservice struct {
	//db ports.PDFPort
}

func New() *PDFservice {
	return &PDFservice{}
}

type InvItems struct {
	Number int
	Item   string
	Amount float64
}

func (s *PDFservice) pdfHeader(pdf *gofpdf.Fpdf, invoice queries.GetInvoiceDataByInvoiceIDRow) *gofpdf.Fpdf {

	pdf.SetFont("Arial", "B", 16)

	// Set the path to the logo image
	logoPath := invoice.LogoUrl

	// Add the logo to the PDF and get its dimensions
	info := pdf.RegisterImageOptions(logoPath, gofpdf.ImageOptions{ReadDpi: true})

	// Set the desired width for the logo
	logoWidth := 30.0 // Adjust this value as needed

	// Calculate the height while preserving the aspect ratio
	logoHeight := info.Height() * logoWidth / info.Width()

	// Position the logo on the right side
	marginRight := 10.0
	pageWidth, _ := pdf.GetPageSize()
	logoX := pageWidth - logoWidth - marginRight
	logoY := 10.0 // Adjust as needed

	// Add the logo to the PDF
	pdf.ImageOptions(logoPath, logoX, logoY, logoWidth, logoHeight, false, gofpdf.ImageOptions{ReadDpi: true}, 0, "")

	// Set font for the text
	pdf.SetFont("Arial", "B", 12)

	// Calculate the position for the text below the logo
	textX := logoX
	textY := logoY + 25 // Adjust spacing between logo and text as needed

	// Add the text below the logo
	pdf.SetXY(textX, textY)
	pdf.CellFormat(logoWidth, 10, invoice.CompanyName, "", 0, "C", false, 0, "")

	if pdf.Err() {
		log.Fatalf("Failed creating PDF report: %s\n", pdf.Error())
	}

	return pdf
}

func formatDate(date pgtype.Date) string {
	dateTime := date.Time

	formattedDate := dateTime.Format("02-Jan-2006")
	return formattedDate
}

func (s *PDFservice) pdfBilling(pdf *gofpdf.Fpdf, invoice queries.GetInvoiceDataByInvoiceIDRow) *gofpdf.Fpdf {
	pdf.Ln(20)

	// Set font for the "INVOICE" title to bold
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(20, 10, "INVOICE")

	// Define billing details
	billingDetails1 := map[string]string{
		"Invoice Number": invoice.InvoiceID.String,
		"Account Number": invoice.AccountNumber,
		"Type":           invoice.InvoiceType.String,
	}
	billingDetails2 := map[string]string{
		"Billing Date": formatDate(invoice.InvoiceDate),
		"Due Date":     formatDate(invoice.InvoiceDueDate),
		"Address":      invoice.Address,
	}

	// Calculate positions for the two columns
	marginLeft := 20.0
	marginTop := 70.0
	columnWidth := 90.0
	rowHeight := 10.0

	// Add billing details in the first column
	currentY := marginTop
	for label, value := range billingDetails1 {
		// Bold for labels
		pdf.SetFont("Arial", "B", 10)
		pdf.SetXY(marginLeft, currentY)
		pdf.CellFormat(columnWidth, rowHeight, fmt.Sprintf("%s:", label), "", 0, "L", false, 0, "")

		// Normal font for values
		pdf.SetFont("Arial", "", 10)
		pdf.SetXY(marginLeft+columnWidth/2, currentY)
		pdf.CellFormat(columnWidth/2, rowHeight, value, "", 0, "L", false, 0, "")

		currentY += rowHeight
	}

	// Add billing details in the second column
	currentY = marginTop
	for label, value := range billingDetails2 {
		// Bold for labels
		pdf.SetFont("Arial", "B", 12)
		pdf.SetXY(marginLeft+columnWidth+5, currentY) // Adjust spacing between columns as needed
		pdf.CellFormat(columnWidth, rowHeight, fmt.Sprintf("%s:", label), "", 0, "L", false, 0, "")

		// Normal font for values
		pdf.SetFont("Arial", "", 10)

		// Check if the label is "Address" and use MultiCell for value if it is
		if label == "Address" {
			pdf.SetXY(marginLeft+columnWidth+5+columnWidth/2, currentY)
			pdf.MultiCell(columnWidth/2, rowHeight, value, "", "L", false)
		} else {
			pdf.SetXY(marginLeft+columnWidth+5+columnWidth/2, currentY)
			pdf.CellFormat(columnWidth/2, rowHeight, value, "", 0, "L", false, 0, "")
		}

		currentY += rowHeight
	}

	return pdf
}

func (s *PDFservice) pdfBillingItems(pdf *gofpdf.Fpdf,
	invoiceItems []queries.GetInvoiceItemDataByInvoiceIDRow) *gofpdf.Fpdf {
	pdf.Ln(10)

	// Set font for the table
	pdf.SetFont("Arial", "B", 12)

	// Define left margin and available width
	marginLeft := 20.0
	marginRight := 20.0
	pageWidth, _ := pdf.GetPageSize()
	availableWidth := pageWidth - marginLeft - marginRight

	// Define column widths
	numberWidth := 20.0 // Adjust as needed
	nameWidth := availableWidth * 0.6
	priceWidth := availableWidth * 0.4

	// Create table header
	header := []string{"_", "Name", "Amount"}

	// Add header row to the table
	pdf.SetFillColor(240, 240, 240)
	pdf.SetTextColor(0, 0, 0)
	for i, str := range header {
		var width float64
		switch i {
		case 0:
			width = numberWidth
		case 1:
			width = nameWidth
		case 2:
			width = priceWidth
		}
		pdf.CellFormat(width, 10, str, "1", 0, "C", true, 0, "")
	}
	pdf.Ln(-1)

	var items []InvItems

	// Define billing items
	for indx, item := range invoiceItems {
		amountFloat := utils.ConvertNumericToFloat64(item.Amount)
		invItem := InvItems{
			Number: indx + 1,
			Item:   item.Item,
			Amount: amountFloat,
		}
		items = append(items, invItem)
	}

	// Set font for the data rows
	pdf.SetFont("Arial", "", 12)

	// Add rows for billing items
	for _, row := range items {
		// Convert each field to a string and add it to the PDF
		pdf.CellFormat(numberWidth, 10, fmt.Sprintf("%d", row.Number), "1", 0, "C", false, 0, "")
		pdf.CellFormat(nameWidth, 10, row.Item, "1", 0, "C", false, 0, "")
		pdf.CellFormat(priceWidth, 10, fmt.Sprintf("%.2f", row.Amount), "1", 0, "C", false, 0, "")
		pdf.Ln(-1)
	}

	// Add total row
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(numberWidth+nameWidth, 10, "Total:", "1", 0, "R", false, 0, "")

	// Calculate the total amount
	totalAmount := 0.0
	for _, item := range items {
		totalAmount += item.Amount
	}
	pdf.CellFormat(priceWidth, 10, fmt.Sprintf("/= %.2f", totalAmount), "1", 0, "C", false, 0, "")

	return pdf
}

// func (s *PDFservice) pdfBillingItems(pdf *gofpdf.Fpdf,
// 	invoiceItems []queries.GetInvoiceItemDataByInvoiceIDRow) *gofpdf.Fpdf {
// 	pdf.Ln(10)

// 	// Set font for the table
// 	pdf.SetFont("Arial", "B", 12)

// 	// Define left margin and available width
// 	marginLeft := 20.0
// 	marginRight := 20.0
// 	pageWidth, _ := pdf.GetPageSize()
// 	availableWidth := pageWidth - marginLeft - marginRight

// 	// Define column widths
// 	numberWidth := 20.0 // Adjust as needed
// 	nameWidth := availableWidth * 0.6
// 	priceWidth := availableWidth * 0.4

// 	// Create table header
// 	header := []string{"_", "Name", "Amount"}

// 	// Add header row to the table
// 	pdf.SetFillColor(240, 240, 240)
// 	pdf.SetTextColor(0, 0, 0)
// 	for i, str := range header {
// 		var width float64
// 		switch i {
// 		case 0:
// 			width = numberWidth
// 		case 1:
// 			width = nameWidth
// 		case 2:
// 			width = priceWidth
// 		}
// 		pdf.CellFormat(width, 10, str, "1", 0, "C", true, 0, "")
// 	}
// 	pdf.Ln(-1)

// 	var items []InvItems

// 	// Define billing items
// 	for indx, item := range invoiceItems {
// 		amountFloat := utils.ConvertNumericToFloat64(item.Amount)
// 		invItem := InvItems{
// 			Number: indx + 1,
// 			Item:   item.Item,
// 			Amount: amountFloat,
// 		}
// 		items = append(items, invItem)

// 	}

// 	// Set font for the data rows
// 	pdf.SetFont("Arial", "", 12)

// 	// Add rows for billing items
// 	for _, row := range items {
// 		for j, str := range row {
// 			var width float64
// 			switch j {
// 			case 0:
// 				width = numberWidth
// 			case 1:
// 				width = nameWidth
// 			case 2:
// 				width = priceWidth
// 			}
// 			pdf.CellFormat(width, 10, str, "1", 0, "C", false, 0, "")
// 		}
// 		pdf.Ln(-1)
// 	}

// 	// Add total row
// 	pdf.SetFont("Arial", "B", 12)
// 	pdf.CellFormat(numberWidth+nameWidth, 10, "Total:", "1", 0, "R", false, 0, "")
// 	pdf.CellFormat(priceWidth, 10, "$60", "1", 0, "C", false, 0, "")

// 	return pdf
// }

func (s *PDFservice) pdfFooter(pdf *gofpdf.Fpdf) *gofpdf.Fpdf {
	// Set font for the footer
	pdf.SetFont("Arial", "", 10)

	// Get page width and height
	pageWidth, pageHeight := pdf.GetPageSize()

	// Calculate Y position for the footer (bottom margin)
	marginBottom := 31.0
	footerY := pageHeight - marginBottom

	// Message for the footer
	footerText := "For official use only - Generated by System"

	// Get width of the footer text
	footerWidth := pdf.GetStringWidth(footerText)

	// Calculate X position for centering the footer
	footerX := (pageWidth - footerWidth) / 2

	// Add footer to the PDF
	pdf.SetXY(footerX, footerY)
	pdf.Cell(0, 10, footerText)

	return pdf
}

func (s *PDFservice) GenerateInvoicePDF(
	invoice queries.GetInvoiceDataByInvoiceIDRow,
	invoiceItems []queries.GetInvoiceItemDataByInvoiceIDRow) error {

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	//Header Fn
	pdf = s.pdfHeader(pdf, invoice)

	//Billing Details
	pdf = s.pdfBilling(pdf, invoice)

	//Billimg Items
	pdf = s.pdfBillingItems(pdf, invoiceItems)

	//Footer
	pdf = s.pdfFooter(pdf)

	err := pdf.OutputFileAndClose(fmt.Sprintf("%s.pdf", invoice.InvoiceID.String))
	if err != nil {
		fmt.Println(err)
	}

	return err
}
