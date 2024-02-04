MiBlog
================================
Simple Blog created using Golang Programming

### Run Application
```shell
docker-compose up -d --build
go run main.go runserver
```

### Migration
```shell
# migrate
go run main.go migrate -e up
go run main.go migrate -e up 0001
go run main.go migrate -e down
go run main.go migrate -e down 0001
go run main.go migrate -e undo

# makemigrations
go run main.go makemigrations -f "create_new_migration_file"
```