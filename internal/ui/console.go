package ui

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

var (
	Cyan  = color.New(color.FgCyan).SprintFunc()
	Green = color.New(color.FgGreen).SprintFunc()
	Red   = color.New(color.FgRed).SprintFunc()
)

func PrintInfo(msg string) {
	fmt.Println(Cyan("ℹ️  " + msg))
}

func PrintSuccess(msg string) {
	fmt.Println(Green("✅ " + msg))
}

func PrintError(msg string) {
	fmt.Fprintln(os.Stderr, Red("❌ "+msg))
}

func NewSpinner(msg string) *spinner.Spinner {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Suffix = " " + msg
	return s
}

func PrintTree(root string, prefix string) {
	files, _ := os.ReadDir(root)
	for i, f := range files {
		connector := "├──"
		if i == len(files)-1 {
			connector = "└──"
		}

		fmt.Printf("%s%s %s\n", prefix, connector, f.Name())

		if f.IsDir() {
			PrintTree(filepath.Join(root, f.Name()), prefix+"│   ")
		}
	}
}
