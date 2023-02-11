package resolver

import (
	"context"

	"github.com/nitintf/graph-go/graph/models"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Impl ResolverImpl
}

type ResolverImpl interface {
	Users(ctx context.Context) (*models.UsersPayload, error)
	CreateUser(ctx context.Context, input models.UserInput) (*models.User, error)
	UpdateUser(ctx context.Context, input models.UserInput) (*models.User, error)
	Me(ctx context.Context) (*models.User, error)
}
