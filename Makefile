# host ?= 103.93.58.89
host ?= 101.255.3.67
port ?= 5432
user ?= postgres
pass ?= 1
db ?= event_organizer

migrate\:init:
	PGPASSWORD=$(pass) psql -U$(user) -h $(host) -d postgres -p $(port) -c "create database $(db);"

migrate\:drop:
	PGPASSWORD=$(pass) psql -U$(user) -h $(host) -d postgres -p $(port) -c "drop database if exists $(db) with (force);"

migrate\:up:
	migrate -database postgresql://$(user):$(pass)@$(host):$(port)/$(db)?sslmode=disable -path migrations up $(version)

migrate\:down:
	migrate -database postgresql://$(user):$(pass)@$(host):$(port)/$(db)?sslmode=disable -path migrations down $(version)

migrate\:reset: 
	$(MAKE) migrate:drop user=$(user) db=$(db)
	$(MAKE) migrate:init user=$(user) db=$(db)
	$(MAKE) migrate:up user=$(user) pass=$(pass) db=$(db)


migrate\:create:
	migrate create -ext sql -dir migrations -seq create_$(table)_table

migrate\:alter:
	migrate create -ext sql -dir migrations -seq alter_$(table)_table

migrate\:insert:
	migrate create -ext sql -dir migrations -seq insert_$(table)_table

