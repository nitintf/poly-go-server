package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"poly-go-server/graph/backend"
	"poly-go-server/graph/directives"
	"poly-go-server/graph/generated"
	"poly-go-server/graph/resolver"
	"poly-go-server/internal/config"
	"poly-go-server/internal/db"
	"poly-go-server/internal/logging"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	ctx := context.Background()

	cfg, err := config.New()

	if err != nil {
		log.Fatal().Err(err).Msg("Init config failed")
	}

	logger := logging.New(cfg)
	ctx = logger.WithContext(ctx)

	dbClient := db.New(cfg, ctx)

	// The server should never exit unless something is wrong.
	logger.Info().Str("ListenAdress", cfg.Port).Msgf("Starting %s service", cfg.ServiceName)
	err = http.ListenAndServe(":"+cfg.Port, getMux(logger, dbClient))

	if err != nil {
		logger.Fatal().Err(err).Msg("Server exit")
	}
}

func getMux(l *zerolog.Logger, dbClient *db.Client) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(rw http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(rw, "Hello!")
	})

	mux.HandleFunc("/graphql", playground.Handler("Graphql Playground", "/query"))

	gqlgenConfig := generated.Config{
		Resolvers: &resolver.Resolver{
			Impl: backend.InitResolvers(l, dbClient),
		},
	}
	gqlgenConfig.Directives.Binding = directives.Binding
	srv := newGqlgenServer(generated.NewExecutableSchema(gqlgenConfig), l)
	mux.Handle("/query", srv)

	// Add CORS options
	corsOptions := cors.Options{
		AllowedOrigins: []string{
			"http://localhost:*",
		},
		AllowCredentials: true,
		MaxAge:           60,
	}

	return cors.New(corsOptions).Handler(mux)
}

func newGqlgenServer(es graphql.ExecutableSchema, l *zerolog.Logger) *handler.Server {
	srv := handler.New(es)

	srv.AroundResponses(func(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
		oc := graphql.GetOperationContext(ctx)
		res := next(ctx)

		startTime := graphql.GetStartTime(ctx)
		latency := time.Until(startTime)

		l.Info().
			Str("operation", oc.OperationName).
			Interface("variables", oc.Variables).
			Str("latency", strings.Split(latency.String(), "-")[1]).
			Interface("errors", res.Errors).
			Interface("data", res.Data).
			Msg("graphql response")

		return res
	})

	srv.AddTransport(transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				origin := r.Header.Get("Origin")
				l.Debug().Str("ComponentName", "GqlServer").
					Interface("body", r.Body).
					Interface("headers", r.Header).
					Msg("Checking origin")
				return strings.HasPrefix(origin, "http://localhost:")
			},
			EnableCompression: true,
		},
		PingPongInterval: 5 * time.Second,
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.SetQueryCache(lru.New(1000))
	srv.Use(extension.Introspection{})
	return srv
}
