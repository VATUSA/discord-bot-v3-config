package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type ConfigFile struct {
	FileName string
	Contents string
	Lines    []string
}

func main() {
	err := generateConfigMap()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
}
func generateConfigMap() error {
	rootPath := "config"
	pfx := len(rootPath) + 1
	tpl := template.Must(template.ParseFiles("generator/templates/configmap.yaml"))

	var configFiles []ConfigFile

	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, e1 error) error {
		if !info.IsDir() && strings.HasSuffix(path, ".yaml") {
			if e1 != nil {
				return e1
			}

			file, err := os.Open(path)
			if err != nil {
				return err
			}
			fileScanner := bufio.NewScanner(file)
			fileScanner.Split(bufio.ScanLines)

			var lines []string

			for fileScanner.Scan() {
				lines = append(lines, fileScanner.Text())
			}

			name := path[pfx:]

			configFiles = append(configFiles, ConfigFile{
				FileName: name,
				Lines:    lines,
			})
		}
		return nil
	})
	if err != nil {
		return err
	}

	f, err := os.OpenFile("manifest/configmap.yaml", os.O_WRONLY+os.O_TRUNC, os.ModeExclusive)
	err = tpl.ExecuteTemplate(f, "configmap.yaml", configFiles)
	if err != nil {
		return err
	}
	return nil
}
