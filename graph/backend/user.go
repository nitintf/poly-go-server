package backend

import (
	"context"
	"errors"

	"github.com/nitintf/graph-go/graph/models"
)

func (r *Resolver) Users(ctx context.Context) (*models.UsersPayload, error) {
	r.logger.Info().Msg("Listin all users - start")

	// users := r.db.Repos.UsersRepo.

	// var queryMods []qm.QueryMod

	// users, count, err := daos.FindAllUsersWithCount(queryMods, ctx)

	// return &graphmodels.UsersPayload{Total: int(count), Users: utils.UsersToGraphQlUsers(users)}, err
	return nil, nil
}

func (r *Resolver) CreateUser(ctx context.Context, input models.UserInput) (*models.User, error) {
	_, err := r.db.UsersRepo.GetUserByEmail(input.Email)

	if err == nil {
		return nil, errors.New("email already exists")
	}

	_, err = r.db.UsersRepo.GetUserByUsername(input.UserName)
	if err == nil {
		return nil, errors.New("username already in used")
	}

	user := &models.User{
		Username:  input.UserName,
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}

	tx, err := r.db.UsersRepo.DB.Begin()

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

func (r *Resolver) UpdateUser(ctx context.Context, input models.UserInput) (*models.User, error) {
	r.logger.Info().Msg("Listin all users - start")
	return nil, nil
}

func (r *Resolver) Me(ctx context.Context) (*models.User, error) {
	r.logger.Info().Msg("Listin all users - start")
	return nil, nil
}
