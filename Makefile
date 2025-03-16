.PHONY:dbup
dbup:
	docker run --rm --name sqltestpgx -e "POSTGRES_PASSWORD=postgres" -p 5432 -d postgres:13-alpine
.PHONY:dbhost
dbhost:
	@docker port sqltestpgx 5432
.PHONY:test
test:
	DATABASE_URL="postgresql://postgres:postgres@$$($(MAKE) dbhost)/postgres" go test -count=1 -race ./...
