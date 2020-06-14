init:
	docker network create gql-demo-network
	docker volume create gql-demo-data

up:
	docker-compose up -d

down:
	docker-compose down

serve:
	docker-compose exec app go run server.go

sh:
	docker-compose exec app sh

psql:
	docker-compose exec db psql -d postgres -U postgres

gen:
	go run github.com/99designs/gqlgen generate
	# go get github.com/vektah/dataloaden

migrate:
	sql-migrate up

dblog:
	docker-compose logs -f db
