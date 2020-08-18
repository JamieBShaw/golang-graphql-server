migrate-create-users:
	migrate create -ext sql -dir postgres/migrations create_users
migrate-create-meetups:
	migrate create -ext sql -dir postgres/migrations create_meetups
migrate-up:
	migrate -path "postgres/migrations" -database "postgres://postgres:postgres@localhost:5432/meetmeup_dev?sslmode=disable" up
migrate-down:
	migrate -path "postgres/migrations" -database "postgres://postgres:postgres@localhost:5432/meetmeup_dev?sslmode=disable" down
gqlgen-gen:
	go run github.com/99designs/gqlgen --verbose
