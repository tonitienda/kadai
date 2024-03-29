start:
    docker-compose up

dev:
    docker-compose -f docker-compose.yaml -f docker-compose.dev.yaml up


start-backend:
    source .env && cd backend-golang-rest/cmd && go run .

unit-test:
    cd backend-golang-rest && GIN_MODE=test go test ./... -v

test-coverage:
    cd backend-golang-rest && GIN_MODE=test go test ./... -cover -coverpkg=./... -count=1 -coverprofile=coverage.out
    cd backend-golang-rest && go tool cover -html=coverage.out -o coverage.html

start-frontend:
    cp .env webapp-nextjs/.env.local
    cd webapp-nextjs && BACKEND_BASE_URL="http://localhost:8080" npm run dev


