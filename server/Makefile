postgresinit:
	sudo docker run --name postgreschat -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15-alpine
postgres:
	sudo docker exec -it postgreschat psql
createdb:
	sudo docker exec -it postgreschat createdb --username=root --owner=root chatify
dropdb:
	sudo docker exec -it postgreschat dropdb chatify
migrateup:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5433/chatify?sslmode=disable" --verbose up 
migratedown:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5433/chatify?sslmode=disable" --verbose down

.PHONY: postgresinit postgres createdb dropdb migrateup migratedown