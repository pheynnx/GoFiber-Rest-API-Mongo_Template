# GoFiber MongoDB Rest API Template

### Demonstrates:
- Full CRUD REST API
- MongoDB native driver; no use of orm/odm
- How to structure a production ready API (Model/Controller/Routes)
- How to implement custom JSON schema validation (database/database.go)
- How to implement custom collection indexes (database/database.go)

### Start:
- Make sure to setup environmental variables before running

### Environmental Variables:
#### File Path: ./config/config.env
- PORT
- MONGO_URI

### Run server:
```bash
go run server.go
```

### Use fresh for hot reload:
```bash
go get github.com/pilu/fresh
fresh
```

### Tools:
- GoFiber v2
- Mongo-driver