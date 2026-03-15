package generator

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

// FileSpec describes one file to be generated.
type FileSpec struct {
	// SrcTemplate is the path to the template file, relative to the templates/ dir.
	SrcTemplate string
	// DstPath is the destination path, relative to the project root.
	// It may itself contain template expressions, e.g. "{{.ServiceName}}/go.mod".
	DstPath   string
	Condition func(c *Config) bool
}

// Generate runs the full project generation for the given config.
func Generate(cfg *Config, templatesDir string) error {
	specs := buildRegistry(cfg)

	for _, spec := range specs {
		if spec.Condition != nil && !spec.Condition(cfg) {
			continue
		}

		// Render destination path (supports template expressions)
		dstPath, err := renderString(spec.DstPath, cfg)
		if err != nil {
			return fmt.Errorf("rendering dst path %q: %w", spec.DstPath, err)
		}

		srcPath := filepath.Join(templatesDir, spec.SrcTemplate)
		if err := renderFile(srcPath, dstPath, cfg); err != nil {
			return fmt.Errorf("rendering %q → %q: %w", spec.SrcTemplate, dstPath, err)
		}
	}

	// Run go mod tidy in the generated project
	projectDir := cfg.ServiceName
	if err := runGoModTidy(projectDir, cfg); err != nil {
		// non-fatal — just warn
		fmt.Printf("⚠️  go mod tidy failed: %v\n", err)
	}

	return nil
}

// renderFile reads a template file, renders it with cfg, and writes to dst.
func renderFile(srcPath, dstPath string, cfg *Config) error {
	contents, err := os.ReadFile(srcPath)
	if err != nil {
		return fmt.Errorf("reading template %q: %w", srcPath, err)
	}

	rendered, err := renderBytes(contents, cfg)
	if err != nil {
		return fmt.Errorf("rendering template %q: %w", srcPath, err)
	}

	if err := os.MkdirAll(filepath.Dir(dstPath), 0755); err != nil {
		return err
	}
	return os.WriteFile(dstPath, rendered, 0644)
}

// renderBytes executes a Go template with cfg as data.
func renderBytes(src []byte, cfg *Config) ([]byte, error) {
	tmpl, err := template.New("").
		Delims("[[", "]]"). // Use [[ ]] to avoid conflicts with Go generics / angle brackets
		Funcs(templateFuncs()).
		Parse(string(src))
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, cfg); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func renderString(s string, cfg *Config) (string, error) {
	b, err := renderBytes([]byte(s), cfg)
	return string(b), err
}

// templateFuncs returns custom template helper functions.
func templateFuncs() template.FuncMap {
	return template.FuncMap{
		"lower":   strings.ToLower,
		"upper":   strings.ToUpper,
		"title":   strings.Title, //nolint:staticcheck
		"replace": strings.ReplaceAll,
	}
}

// runGoModTidy generates the go.mod and runs `go mod tidy` in the new project.
func runGoModTidy(projectDir string, cfg *Config) error {
	goModContent := buildGoMod(cfg)
	goModPath := filepath.Join(projectDir, "go.mod")
	if err := os.WriteFile(goModPath, []byte(goModContent), 0644); err != nil {
		return err
	}

	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = projectDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
