# learn_app
## run db postgresql:
- $docker-compose up -d
## run migrate
- $migrate -path ./schema -database 'postgres://postgres:toor-555@localhost:5432/postgres?sslmode=disable' up
## run app:
- $go run cmd/main/app.go
