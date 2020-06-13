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

migrate:
	sql-migrate up
