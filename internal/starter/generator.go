package starter

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func Generate(name, framework, db, log string) error {
	if framework != "echo" {
		return fmt.Errorf("framework %s is not yet supported, only echo is supported for now", framework)
	}

	src := filepath.Join("templates", framework)
	dst := filepath.Join(".", name)

	data := map[string]string{
		"AppName":   name,
		"DB":        db,
		"Logger":    log,
		"Framework": framework,
	}

	return copyAndRender(src, dst, data)
}

func copyAndRender(srcDir, dstDir string, data map[string]string) error {
	return filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath := strings.TrimPrefix(path, srcDir)
		relPath = strings.TrimPrefix(relPath, string(filepath.Separator))
		dstPath := filepath.Join(dstDir, relPath)

		if info.IsDir() {
			return os.MkdirAll(dstPath, os.ModePerm)
		}

		contents, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		tmpl, err := template.New(info.Name()).Parse(string(contents))
		if err != nil {
			return err
		}

		f, err := os.Create(dstPath)
		if err != nil {
			return err
		}
		defer f.Close()

		return tmpl.Execute(f, data)
	})
}
