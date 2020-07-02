# Template Golang Rest API (ready for full stack)

## Start:

- Setup environmental variables

```bash
'go run server.go
```

## Using 'air' for live reload:

```bash
go get -u github.com/cosmtrek/air
air
```

## Tools:

- Fiber
- Gorm

## Environmental Variables:

- ./config/config.env
- PORT=
- DATABASE=

## Frontend (full stack):

- uncomment //app.Static("/", "PATH to build") in server.go (line 39)
- adjusted path for your build location
