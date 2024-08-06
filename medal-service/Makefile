DB_URL = postgres://postgres:1@localhost:5432/medalsdb?sslmode=disable

gen-proto:
	@protoc \
	--go_out=. \
	--go-grpc_out=. \
	protos/medal.proto

migrate-up:
	migrate -path ./db/migrations -database ${DB_URL} up

migrate-down:
	migrate -path ./db/migrations -database ${DB_URL} down

migrate-force:
	migrate -path=db/migrations -database ${DB_URL} verbose -force 1
