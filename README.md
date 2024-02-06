MiBlog
================================
Simple Blog created using Golang Programming

### Run Application
```shell
# running database with docker-compose
docker-compose up -d --build

# run application
go run main.go runserver
# or
make runserver
```

* Access to: http://localhost:8000
* Documentation: http://localhost:8000/docs/index.html



### Docs
```shell
# Install Swaggo
go install github.com/swaggo/swag/cmd/swag@latest

# Generate Docs
swag init --parseDependency --parseInternal
# or
make docs
```


### Migration
```shell
# migrate
go run main.go migrate -e up
# or
make migrate

# Others
go run main.go migrate -e up 0001
go run main.go migrate -e down
go run main.go migrate -e down 0001
go run main.go migrate -e undo


# makemigrations
go run main.go makemigrations -f "create_new_migration_file"
# or
make makemigrations name="create_new_migration_file"
```

### Contributors
<table>
  <tr>
    <td align="center">
      <a href="https://www.linkedin.com/in/agung96tm/">
        <img src="https://avatars.githubusercontent.com/u/1901484?v=4" width="100px;" alt=""/><br />
        <b>Agung Yuliyanto</b><br>
      </a>
      <div>💻</div>
    </td>
  </tr>
</table>