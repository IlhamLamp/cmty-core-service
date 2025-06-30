```
cmty-project-service
├─ .air.toml
├─ cmd
│  ├─ main.go
│  └─ repos.go
├─ config
│  └─ config.go
├─ controllers
│  ├─ member_controller.go
│  └─ project_controller.go
├─ database
│  ├─ connection.go
│  ├─ migrations
│  │  ├─ 001_create_projects.down.sql
│  │  ├─ 001_create_projects.up.sql
│  │  ├─ 002_create_members.down.sql
│  │  └─ 002_create_members.up.sql
│  └─ seeders
│     ├─ 01_project.json
│     └─ 02_member.json
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
│  ├─ member_repository.go
│  └─ project_repository.go
├─ routes
│  └─ router.go
├─ services
│  ├─ member_service.go
│  └─ project_service.go
├─ types
│  └─ config_type.go
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
