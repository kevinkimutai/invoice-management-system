package app

import (
	"github.com/kevinkimutai/invoice-management-system/internal/domain"
	"github.com/kevinkimutai/invoice-management-system/internal/ports"
	"github.com/kevinkimutai/invoice-management-system/internal/utils"
)

type UserRepo struct {
	db ports.UserRepoPort
}

func NewUserRepo(db ports.UserRepoPort) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) GetAllUsers(params domain.Params) (domain.UsersFetch, error) {
	limitParams := utils.GetParams(params)

	usersData, err := r.db.GetUsers(limitParams)

	return usersData, err
}
