.PHONY: gqlgen
gqlgen:         ## Generate gqlgen from schema
	@go run github.com/99designs/gqlgen generate
