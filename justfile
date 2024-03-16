start-backend:
    cd backend && go run cmd/main.go

unit-test:
    cd backend-golang-rest && go test ./...
