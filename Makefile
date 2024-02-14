composeupdb:
	docker compose -f docker-compose.db.yaml up -d

composedowndb:
	docker compose -f "docker-compose.db.yaml" down

composeupimage:
	docker compose -f docker-compose.image.yaml up -d

composedownimage:
	docker compose -f "docker-compose.image.yaml" down

createuserdb:
	docker exec -it user-db createdb --username=root --owner=root user_db

dropuserdb:
	docker exec -it user-db dropdb --username=root user_db

creategroupdb:
	docker exec -it group-db createdb --username=root --owner=root group_db

dropgroupdb:
	docker exec -it group-db dropdb --username=root group_db

createusermigration:
	migrate create -ext sql -dir pkg/database/postgres/userdb/migrations -seq init_user_schema

creategroupmigration:
	migrate create -ext sql -dir pkg/database/postgres/groupdb/migrations -seq init_group_schema

migrateuserup:
	migrate -path pkg/database/postgres/userdb/migrations -database "postgresql://root:123456@localhost:5432/user_db?sslmode=disable" -verbose up

migrateuserdown:
	migrate -path pkg/database/postgres/userdb/migrations -database "postgresql://root:123456@localhost:5432/user_db?sslmode=disable" -verbose down

migrategroupup:
	migrate -path pkg/database/postgres/groupdb/migrations -database "postgresql://root:123456@localhost:5433/group_db?sslmode=disable" -verbose up

migrategroupdown:
	migrate -path pkg/database/postgres/groupdb/migrations -database "postgresql://root:123456@localhost:5433/group_db?sslmode=disable" -verbose down

runauth:
	go run main.go ./env/dev/.env.auth

runuser:
	go run main.go ./env/dev/.env.user

.PHONY: composeupdb composedowndb sqlc migrateuserdown migrateuserup createuserdb dropuserdb createusermigration runauth runuser composeupdb composedowndb