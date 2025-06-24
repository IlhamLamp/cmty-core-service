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
│  │  ├─ 001_create_projects.up.sql
│  │  └─ 001_delete_projects.down.sql
│  └─ seeders
│     └─ 01_project.json
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
└─ types
   └─ config_type.go

```

1. Menjalankan Migration

```
migrate -path ./database/migrations -database "postgres://postgres:postgres@localhost:5432/projects?sslmode=disable" up
```

2. Rollback Migration
   Jika ingin rollback satu step:

```
migrate -path ./database/migrations -database "postgres://postgres:postgres@localhost:5432/projects?sslmode=disable" down 1
```
