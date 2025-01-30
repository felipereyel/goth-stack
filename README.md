# goth oidc

## Components
- Go + Fiber
- Tailwind
- HTMX
- SQLite
- Oidc auth

## Extenal dependencies

### Templ
`templ` is used to generate templates in go format. Install the CLI with:

```bash
    go install github.com/a-h/templ/cmd/templ@latest
```

### air (optional)
`air` is used to hot reload the application. Install the CLI with:
```bash
    go install github.com/cosmtrek/air@latest
```

### Tailwind and Htmx
This project download the statics to embed in the bundle

## How to run
Before running be sure to add all required environment variables (see [env example](.env.example)) 

- make: `make statics` to download statics
- serve: `go run main.go serve` or `air serve`
- migrate up: `go run main.go migrate:up` or `air migrate:up`
- migrate down: `go run main.go migrate:down N` or `air migrate:down N` where N is the number of migrations down

## Roadmap

### Done
- Go
- HTMx
- Tailwind
- Templ
- Custom auth
- JWT cookie auth
- Embeded Migrations

### Next
- [ ] Access + refresh tokens
- [ ] SQL Autogeneration with [sqlc](https://github.com/sqlc-dev/sqlc)
