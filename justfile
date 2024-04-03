start-go:
    COMPOSE_PROJECT_NAME="kadai-nextjs-go" docker-compose -f docker-compose.yaml -f compose.next-frontend.yaml -f compose.go-backend.yaml up --build

start-js:
    COMPOSE_PROJECT_NAME="kadai-nextjs-js" docker-compose -f docker-compose.yaml -f compose.next-frontend.yaml -f compose.js-backend.yaml up --build

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


