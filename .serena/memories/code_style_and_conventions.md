# go-start Code Style and Conventions

## Language
- Go 1.23+
- Standard Go formatting (gofmt/goimports)

## Naming Conventions
- **Packages**: lowercase, short, single-word (e.g., `generator`, `ui`)
- **Types/Structs**: PascalCase (e.g., `Config`, `FileSpec`)
- **Functions**: PascalCase for exported (e.g., `Generate`), camelCase for unexported (e.g., `renderFile`, `buildRegistry`)
- **Variables**: camelCase
- **Constants**: PascalCase or ALL_CAPS for true constants
- **Interface methods**: PascalCase

## Error Handling
- Use `fmt.Errorf("context: %w", err)` for wrapping errors
- Return errors from functions; handle at call site
- Non-fatal errors: print a warning with `fmt.Printf("⚠️  <message>: %v\n", err)` and continue
- Fatal errors: `os.Exit(1)` in CLI command handlers, or use `exitOnErr()` helper in cmd/

## Code Organization Patterns
- The `Config` struct is the central data object — passed by pointer throughout
- `FileSpec` pattern: each generated file has `{SrcTemplate, DstPath, Condition}` — Condition is nil (always included) or a func
- Template conditions use helper predicates: `isEcho(cfg)`, `isGin(cfg)`, `isFiber(cfg)`, `cfg.HasPostgres`, etc.

## Template Files (.tmpl)
- Go `text/template` syntax (`{{ }}` delimiters)
- Template data is the `*Config` struct (accessed as `.ServiceName`, `.ModuleName`, etc.)
- Use custom template functions from `templateFuncs()` (e.g., `lower`, `title`)
- Template files use standard Go formatting within them (they generate valid Go code)

## Comments
- Exported symbols should have GoDoc comments
- Inline comments for non-obvious logic
- No excessive commenting — code should be self-explanatory

## Module/Import Style
- Standard library imports first
- Third-party imports second (separated by blank line)
- Internal package imports last
- Use `goimports` ordering
