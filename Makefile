createmigration:
	migrate create -ext=sql -dir sql/migrations -seq init

migrate:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose up

migratedown:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose down

up:
	docker-compose up

init:
	make migrate && go run ./cmd/ordersystem/main.go ./cmd/ordersystem/wire_gen.go

.PHONY: migrate migratedown createmigration build init