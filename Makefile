run:
	@go run main.go

air: 
	@air

db/create: 
	@docker compose up db --build

db/start:
	@docker compose up db

templ:
	@templ generate -watch --proxy="http://localhost:3000"

dev:
	make -j3 db/start air templ

seed:
	@go run ./db/models/seed.go  