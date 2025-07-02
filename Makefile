DB_HOST?=localhost
DB_USER?=postgres
DB_PASSWORD?=123
DB_PORT?=5434
DB_NAME?=movietar
DB_SSLMODE?=disable

MIGRATION_DIR?=./migrations

MIGRATE=migrate -source file://$(MIGRATION_DIR) -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)"

migration_create:
	migrate create -seq -dir $(MIGRATION_DIR) -ext sql $(name)

migration_up: 
	$(MIGRATE) up

migration_down: 
	$(MIGRATE) down 1

migration_drop: 
	$(MIGRATE) drop -f

build_win:
	GOOS=windows GOARC=amd64 go build -o build/app.exe main.go
build_linux: