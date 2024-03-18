start-backend:
    cd backend-golang-rest && go run cmd/main.go

unit-test:
    cd backend-golang-rest && go test ./...
