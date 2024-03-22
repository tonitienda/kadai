start-backend:
    cd backend-golang-rest && go run cmd/main.go

unit-test:
    cd backend-golang-rest && GIN_MODE=test go test ./... -v

test-coverage:
    cd backend-golang-rest && GIN_MODE=test go test ./... -cover -coverpkg=./... -count=1 -coverprofile=coverage.out
    cd backend-golang-rest && go tool cover -html=coverage.out -o coverage.html


