# Suggested Commands for go-start Development

## Running the Tool
```bash
# Run the CLI wizard directly
go run main.go new

# Build the binary
go build -o go-start .

# Install globally (optional)
sudo mv go-start /usr/local/bin/
go-start new
```

## Development
```bash
# Format code
gofmt -w .
# or
goimports -w .

# Vet code
go vet ./...

# Run tests (if any exist)
go test ./...

# Tidy dependencies
go mod tidy
```

## Linting
```bash
# If golangci-lint is installed
golangci-lint run ./...
```

## Git
```bash
git status
git log --oneline -n 10
git diff
```

## macOS-specific utilities
```bash
find . -name "*.tmpl"           # find template files
grep -r "pattern" . --include="*.go"   # search in Go files
ls -la templates/               # list template dirs
```

## Generated Project Commands (for end users)
These are commands in the *generated* project (not this tool):
```bash
make run          # go run ./cmd/http/main.go
make build        # build binary
make test         # go test ./... -v -cover
make lint         # golangci-lint run ./...
make tidy         # go mod tidy
make swagger      # swag init
make migrate-up   # run pending migrations
make migrate-down # rollback last migration
make docker-up    # docker compose up --build -d
make docker-down  # docker compose down
```
