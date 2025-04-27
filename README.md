# goth oidc

## Components
- Go + Pocketbase
- Tailwind
- HTMX

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
- migrate collections: `go run main.go migrate collections` or `air migrate collections`

## Roadmap

### Done
- Go
- HTMx
- Templ
- Tailwind
- Pocketbase
