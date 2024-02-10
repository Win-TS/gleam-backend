## Installation

Database Docker Image
```bash
# composing dbs in docker-compose.db.yaml file
$ make composeup

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