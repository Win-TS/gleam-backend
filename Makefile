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

migrateuserupdocker:
	migrate -path pkg/database/postgres/userdb/migrations -database "postgresql://root:123456@user-db:5432/user_db?sslmode=disable" -verbose up

migrateuserdowndocker:
	migrate -path pkg/database/postgres/userdb/migrations -database "postgresql://root:123456@user-db:5432/user_db?sslmode=disable" -verbose down

migrategroupupdocker:
	migrate -path pkg/database/postgres/groupdb/migrations -database "postgresql://root:123456@group-db:5432/group_db?sslmode=disable" -verbose up

migrategroupdowndocker:
	migrate -path pkg/database/postgres/groupdb/migrations -database "postgresql://root:123456@group-db:5432/group_db?sslmode=disable" -verbose down

runauth:
	go run main.go ./env/dev/.env.auth

runuser:
	go run main.go ./env/dev/.env.user

rungroup:
	go run main.go ./env/dev/.env.group

genauthproto:
	protoc --go_out=. --go_opt=paths=source_relative \
	    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	    ./modules/auth/authPb/authPb.proto

genuserproto:
	protoc --go_out=. --go_opt=paths=source_relative \
	    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	    ./modules/user/userPb/userPb.proto

gengroupproto:
	protoc --go_out=. --go_opt=paths=source_relative \
	    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	    ./modules/group/groupPb/groupPb.proto

backupuserdb:
	docker exec -t user-db pg_dumpall -c -U root > backup/user-backup.sql 

restoreuserdb:
	cat backup/user-backup.sql | docker exec -i user-db psql -U root -d postgres   

backupgroupdb:
	docker exec -t group-db pg_dumpall -c -U root > backup/group-backup.sql 

restoregroupdb:
	cat backup/group-backup.sql | docker exec -i group-db psql -U root -d postgres   

applyauthservice:
	kubectl apply -f ./build/auth/auth-service.yaml

applyuserservice:
	kubectl apply -f ./build/user/user-service.yaml

applygroupservice:
	kubectl apply -f ./build/group/group-service.yaml

createconfigmap:
	kubectl create configmap app-env --from-file=./env/prod/.env

applyingress:
	kubectl apply -f ./build/gleam-api-ingress.yaml

applyauthdeployment:
	kubectl apply -f ./build/auth/auth-deployment.yaml

applyuserdeployment:
	kubectl apply -f ./build/user/user-deployment.yaml

applygroupdeployment:
	kubectl apply -f ./build/group/group-deployment.yaml

.PHONY: composeupdb composedowndb migrateuserdown migrateuserup createuserdb dropuserdb createusermigration runauth runuser composeupdb composedowndb migrateuserupdocker migrateuserdowndocker migrategroupupdocker migrategroupdowndocker genuserproto gengroupproto genauthproto backupuserdb restoreuserdb backupgroupdb restoregroupdb applyauthservice applyuserservice applygroupservice createconfigmap applyingress applyauthdeployment applyuserdeployment applygroupdeployment