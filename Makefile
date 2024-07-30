run:
	go run .

build:
	go build main.go -o rssagg

test:
	go test -v

migrate:
	goose --dir sql/schema postgres "postgres://postgres:postgres@localhost:5432/rssagg" up

rollback:
	goose --dir sql/schema postgres "postgres://postgres:postgres@localhost:5432/rssagg" down

generate:
	sqlc generate
