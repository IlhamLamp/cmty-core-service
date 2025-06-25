```
cmty-project-service
├─ .air.toml
├─ cmd
│  └─ main.go
├─ config
│  └─ config.go
├─ controllers
│  └─ project_controller.go
├─ database
│  ├─ connection.go
│  ├─ migrations
│  │  ├─ 001_create_members.up.sql
│  │  ├─ 001_delete_member.down.sql
│  │  ├─ 002_create_projects.up.sql
│  │  └─ 002_delete_projects.down.sql
│  └─ seeders
│     └─ 02_project.json
├─ go.mod
├─ go.sum
├─ models
│  ├─ address.go
│  ├─ member.go
│  ├─ project.go
│  ├─ role.go
│  └─ tag.go
├─ README.md
├─ repository
│  └─ project_repository.go
├─ routes
│  └─ router.go
├─ services
│  └─ project_service.go
├─ types
│  ├─ config_type.go
│  └─ response.go
└─ utils
   └─ response.go
```

1. Migration Up

```
migrate -path ./database/migrations -database "postgres://username:password@hostname:5432/database?sslmode=disable" up
```

2. Migration Down

```
migrate -path ./database/migrations -database "postgres://username:password@hostname:5432/database?sslmode=disable" down
```

3. Rollback Migration
   ex: 1 step

```
migrate -path ./database/migrations -database "postgres://username:password@hostname:5432/database?sslmode=disable" down 1
```
