package db

import (
	"context"

	"poly-go-server/internal/config"

	"github.com/go-pg/pg/v10"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

type Repos struct {
	UsersRepo
}

type Client struct {
	Client *pg.DB
	*Repos
}

func New(cfg *config.Config, ctx context.Context) *Client {
	opts, err := pg.ParseURL(cfg.LocalDbURL)

	if err != nil {
		log.Fatal().Err(err).Msgf("URL Parse failed")
	}

	db := pg.Connect(opts)

	err = db.Ping(ctx)

	if err != nil {
		log.Fatal().Err(err).Msgf("Database connection failed")
	}

	return initReposWithClient(db)
}

func initReposWithClient(db *pg.DB) *Client {
	// init repos and pass as dependency
	userRepo := UsersRepo{DB: db}

	c := &Client{
		Client: db,
		Repos: &Repos{
			UsersRepo: userRepo,
		},
	}

	return c
}
