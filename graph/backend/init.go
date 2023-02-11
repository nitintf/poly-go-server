package backend

import (
	"github.com/nitintf/graph-go/internal/db"
	"github.com/rs/zerolog"
)

type Resolver struct {
	logger *zerolog.Logger
	db     *db.Client
}

func InitResolvers(l *zerolog.Logger, db *db.Client) *Resolver {
	return &Resolver{
		logger: l,
		db:     db,
	}
}
