package app

import (
	"github.com/kevinkimutai/invoice-management-system/internal/domain"
	"github.com/kevinkimutai/invoice-management-system/internal/ports"
	"github.com/kevinkimutai/invoice-management-system/internal/utils"
)

type CompanyRepo struct {
	db ports.CompanyRepoPort
}

func NewCompanyRepo(db ports.CompanyRepoPort) *CompanyRepo {
	return &CompanyRepo{
		db: db,
	}
}

func (r *CompanyRepo) GetAllCompanies(params domain.Params) (domain.CompaniesFetch, error) {
	limitParams := utils.GetParams(params)

	companies, err := r.db.GetCompanies(limitParams)

	return companies, err
}

func (r *CompanyRepo) CreateCompany(company domain.Company) (domain.Company, error) {
	com, err := r.db.CreateCompany(company)

	return com, err

}
