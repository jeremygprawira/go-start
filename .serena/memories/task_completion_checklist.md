# Task Completion Checklist

After completing any coding task in go-start, verify:

## 1. Correctness
- [ ] New templates are registered in `internal/generator/registry.go` via `buildRegistry()`
- [ ] Template conditions (`Condition func(*Config) bool`) are properly set (nil = always, otherwise use cfg predicates)
- [ ] Template syntax is valid Go `text/template` — test by running `go run main.go new` end-to-end
- [ ] Destination paths in `DstPath` use `{{.ServiceName}}` prefix for the generated project root

## 2. Config Alignment
- [ ] New options added to CLI wizard (`cmd/new.go`) are also added to the `Config` struct (`internal/generator/config.go`)
- [ ] Helper methods on `Config` are updated if new boolean flags are added

## 3. Go Quality
- [ ] Run `go vet ./...` — must pass with no errors
- [ ] Run `go mod tidy` — keep go.mod/go.sum clean
- [ ] Code is formatted with `gofmt -w .`

## 4. Template Quality
- [ ] Generated Go code from templates compiles cleanly
- [ ] Template files follow the same patterns as existing `.tmpl` files
- [ ] Conditional imports in templates use `{{ if .HasPostgres }}` patterns consistently

## 5. Integration
- [ ] Test the full generation flow: `go run main.go new` and select the affected options
- [ ] Verify the generated project has the correct files and content
- [ ] If adding a new framework/logger/db: add templates to all relevant paths (core, delivery, setup, teardown)
