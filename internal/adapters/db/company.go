package db

import (
	"context"
	"math"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kevinkimutai/invoice-management-system/internal/adapters/queries"
	"github.com/kevinkimutai/invoice-management-system/internal/domain"
	"github.com/kevinkimutai/invoice-management-system/internal/utils"
)

func (db *DBAdapter) GetCompanies(params utils.LimitParams) (domain.CompaniesFetch, error) {
	ctx := context.Background()

	//Get Companies
	companies, err := db.queries.ListCompany(ctx, queries.ListCompanyParams(params))
	if err != nil {
		return domain.CompaniesFetch{}, err

	}

	//Get Count
	count, err := db.queries.GetTotalCompaniesCount(ctx)
	if err != nil {
		return domain.CompaniesFetch{}, err

	}

	//Get Page
	page := utils.GetPage(params.Offset, params.Limit)

	//map struct
	var c []domain.Company

	for _, item := range companies {
		company := domain.Company{
			CompanyID:   item.CompanyID,
			LogoUrl:     item.LogoUrl,
			UserID:      item.UserID.Int64,
			CompanyName: item.CompanyName,
			CreatedAt:   item.CreatedAt.Time,
		}
		// Append the struct to the struct array
		c = append(c, company)
	}

	return domain.CompaniesFetch{
		Page:          page,
		NumberOfPages: uint(math.Ceil(float64(count) / float64(params.Limit))),
		Total:         uint(count),
		Data:          c,
	}, nil
}

func (db *DBAdapter) CreateCompany(company domain.Company) (domain.Company, error) {
	ctx := context.Background()

	intString := utils.ConvertInt64ToString(company.UserID)
	var pgtypeint pgtype.Int8

	pgtypeint.Scan(intString)

	companyParams := queries.CreateCompanyParams{
		UserID:      pgtypeint,
		LogoUrl:     company.LogoUrl,
		CompanyName: company.CompanyName,
	}

	com, err := db.queries.CreateCompany(ctx, companyParams)
	if err != nil {
		return domain.Company{}, err
	}

	return domain.Company{
		CompanyID:   com.CompanyID,
		LogoUrl:     com.LogoUrl,
		UserID:      com.UserID.Int64,
		CompanyName: com.CompanyName,
		CreatedAt:   com.CreatedAt.Time,
	}, nil
}
