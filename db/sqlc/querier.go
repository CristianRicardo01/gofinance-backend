// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"context"
)

type Querier interface {
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
	CreateCategory(ctx context.Context, arg CreateCategoryParams) (Category, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteAccout(ctx context.Context, id int32) error
	DeleteCategory(ctx context.Context, id int32) error
	GetAccontGraph(ctx context.Context, arg GetAccontGraphParams) (int64, error)
	GetAccontReports(ctx context.Context, arg GetAccontReportsParams) (int64, error)
	GetAccount(ctx context.Context, id int32) (Account, error)
	GetAccounts(ctx context.Context, arg GetAccountsParams) ([]GetAccountsRow, error)
	GetCategories(ctx context.Context, arg GetCategoriesParams) ([]Category, error)
	GetCategory(ctx context.Context, id int32) (Category, error)
	GetUser(ctx context.Context, username string) (User, error)
	GetUserById(ctx context.Context, id int32) (User, error)
	UpdateAccont(ctx context.Context, arg UpdateAccontParams) (Account, error)
	UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (Category, error)
}

var _ Querier = (*Queries)(nil)
