up:
	docker-compose up -d

down:
	docker-compose down

serve:
	docker-compose exec app go run server.go

sh:
	docker-compose exec app sh
