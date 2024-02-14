## Installation

Install Packages
```bash
$ brew install golang-migrate
$ brew install bufbuild/buf/buf
$ go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

Database Docker Image
```bash
# composing dbs in docker-compose.db.yaml file
$ make composeupdb

# user service database (postgresql)
$ make createuserdb
$ make migrateuserup
```

## Running the app

```bash
# auth service
$ make runauth

# user service
$ make runuser
```

## Database Migrations

```bash
# *user service*
# new migration
$ make createusermigration
# migrate up
$ make migrateuserup
# migratedown
$ make migrateuserdown
```

