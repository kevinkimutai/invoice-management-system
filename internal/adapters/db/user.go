package db

import (
	"context"
	"math"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kevinkimutai/invoice-management-system/internal/adapters/queries"
	"github.com/kevinkimutai/invoice-management-system/internal/domain"
	"github.com/kevinkimutai/invoice-management-system/internal/utils"
)

func (db *DBAdapter) GetUsers(params utils.LimitParams) (domain.UsersFetch, error) {
	ctx := context.Background()

	usersParams := queries.ListUsersParams{
		Limit:  params.Limit,
		Offset: params.Offset,
	}

	//Get Users
	users, err := db.queries.ListUsers(ctx, usersParams)
	if err != nil {
		return domain.UsersFetch{}, err

	}

	//Get Count
	count, err := db.queries.GetTotalUsersCount(ctx)
	if err != nil {
		return domain.UsersFetch{}, err

	}

	//Get Page
	page := utils.GetPage(params.Offset, params.Limit)

	//map struct
	var u []domain.User

	for _, item := range users {
		user := domain.User{
			UserID:    item.UserID,
			Email:     item.Email.String,
			Name:      item.Name.String,
			CreatedAt: item.CreatedAt.Time,
		}
		// Append the struct to the struct array
		u = append(u, user)
	}

	return domain.UsersFetch{
		Page:          page,
		NumberOfPages: uint(math.Ceil(float64(count) / float64(params.Limit))),
		Total:         uint(count),
		Data:          u,
	}, nil
}

func (db *DBAdapter) CreateUser(user queries.CreateUserParams) (queries.User, error) {
	ctx := context.Background()

	u, err := db.queries.CreateUser(ctx, user)
	if err != nil {
		return u, err
	}

	return u, nil
}

func (db *DBAdapter) GetUserByEmail(email string) (queries.User, error) {
	ctx := context.Background()

	var pgtypetext pgtype.Text
	pgtypetext.Scan(email)

	user, err := db.queries.GetUserByEmail(ctx, pgtypetext)
	if err != nil {
		return queries.User{}, err
	}

	return user, nil
}
