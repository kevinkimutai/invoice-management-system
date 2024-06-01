package utils

import (
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kevinkimutai/invoice-management-system/internal/domain"
)

type LimitParams struct {
	Limit  int32
	Offset int32
}

type InvoiceParams struct {
	InvoiceID      pgtype.Text
	UserID         pgtype.Int8
	CompanyID      pgtype.Int8
	Address        string
	AccountNumber  string
	TotalAmount    pgtype.Numeric
	InvoiceDate    pgtype.Date
	InvoiceDueDate pgtype.Date
	Status         string
	InvoiceType    pgtype.Text
	Items          []InvoiceItemParams
}

type InvoiceItemParams struct {
	InvoiceID string
	Item      string
	Amount    pgtype.Numeric
}

var LIMIT, OFFSET int32 = 10, 0

func GetParams(params domain.Params) LimitParams {
	if params.Limit != "" {
		items, _ := strconv.Atoi(params.Limit)

		LIMIT = int32(items)

	}
	if params.Page != "" {
		page, _ := strconv.Atoi(params.Page)

		if page < 1 {
			page = 1
		}

		OFFSET = (int32(page) - 1) * LIMIT

	}

	return LimitParams{
		Limit:  LIMIT,
		Offset: OFFSET,
	}
}

func GetPage(offset, limit int32) uint {
	return uint((offset / limit) + 1)
}

func ConvertInt64ToString(i int64) string {
	str := strconv.FormatInt(i, 10)
	return str
}

func ConvertFloatToNumeric(f float64) pgtype.Numeric {
	var numeric pgtype.Numeric
	//Convert To Str
	str := strconv.FormatFloat(f, 'f', 2, 64)

	numeric.Scan(str)

	return numeric

}

func GetParamsTypes(item domain.InvoiceItem) InvoiceItemParams {
	return InvoiceItemParams{
		InvoiceID: item.InvoiceID,
		Item:      item.Item,
		Amount:    ConvertFloatToNumeric(item.Amount),
	}

}

func ConvertNumericToFloat64(numeric pgtype.Numeric) float64 {
	fval, _ := numeric.Value()

	//Convert To Float64
	var floatVal float64
	if strVal, ok := fval.(string); ok {
		floatVal, _ = strconv.ParseFloat(strVal, 64)
	}
	return floatVal

}

func ConvertInvoiceTypes(domain domain.Invoice) InvoiceParams {
	var invoiceidtext pgtype.Text
	var invoicetypetext pgtype.Text
	var userid pgtype.Int8
	var companyid pgtype.Int8
	var total pgtype.Numeric
	var date pgtype.Date
	var duedate pgtype.Date

	totalstr := strconv.FormatFloat(domain.TotalAmount, 'f', 2, 64)

	invoiceidtext.Scan(domain.InvoiceID)
	userid.Scan(domain.UserID)
	companyid.Scan(domain.CompanyID)
	total.Scan(totalstr)
	invoicetypetext.Scan(domain.InvoiceType)
	date.Scan(domain.InvoiceDate)
	duedate.Scan(domain.InvoiceDueDate)

	var items []InvoiceItemParams

	for _, val := range domain.Items {
		i := GetParamsTypes(val)

		items = append(items, i)
	}

	return InvoiceParams{
		InvoiceID:      invoiceidtext,
		UserID:         userid,
		CompanyID:      companyid,
		Address:        domain.Address,
		AccountNumber:  domain.AccountNumber,
		TotalAmount:    total,
		InvoiceDate:    date,
		InvoiceDueDate: duedate,
		InvoiceType:    invoicetypetext,
		Items:          items,
	}

}
