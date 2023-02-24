package resolver

import (
	"context"

	"poly-go-server/graph/models"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Impl ResolverImpl
}

type ResolverImpl interface {
	Register(ctx context.Context, input models.RegisterUserInput) (*models.TokenResponse, error)
	Login(ctx context.Context, input models.LoginInput) (*models.TokenResponse, error)
	Me(ctx context.Context) (*models.User, error)
}
