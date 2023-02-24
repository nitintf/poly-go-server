package backend

import (
	"context"
	"errors"

	"poly-go-server/graph/models"
)

func (r *Resolver) Register(ctx context.Context, input models.RegisterUserInput) (*models.TokenResponse, error) {
	_, err := r.db.UsersRepo.GetUserByEmail(input.Email)

	if err == nil {
		return nil, errors.New("email already exists")
	}

	user := &models.User{

		Email: input.Email,
	}

	tx, err := r.db.Client.Begin()

	if err != nil {
		r.logger.Fatal().Err(err).Msg("error creating a transaction")
		return nil, errors.New("something went wrong")
	}
	defer tx.Rollback()

	userRes, err := r.db.UsersRepo.CreateUser(tx, user)

	if err != nil {
		r.logger.Fatal().Err(err).Msg("error creating a user")
		return user, err
	}

	if err := tx.Commit(); err != nil {
		r.logger.Fatal().Err(err).Msg("error while committing")
		return nil, err
	}

	return userRes, nil
}

func (r *Resolver) Login(ctx context.Context, input models.LoginInput) (*models.TokenResponse, error) {
	r.logger.Info().Msg("Listin all users - start")
	return nil, nil
}

func (r *Resolver) Me(ctx context.Context) (*models.User, error) {
	r.logger.Info().Msg("Listin all users - start")
	return nil, nil
}
