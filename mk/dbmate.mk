.PHONY: db-up db-new db-down db-dump db-migrate

_start-db:
	@docker-compose up -d db

db-up: _start-db ## create the database (if it does not already exist) and run any pending migrations
	@dbmate up

db-new:         ## generate a new migration file
	@dbmate new $(migration)

db-down: _start-db ## alias for rollback
	@dbmate down

db-dump:        ## Drop the database
	@dbmate dump

db-migrate:     ## run any pending migrations
	@dbmate migrate
