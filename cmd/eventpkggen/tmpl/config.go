package tmpl

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

type TagInfo struct {
	Tag            string
	TagName        string
	TagPackageName string
}
type TemplateStruct struct {
	PackageName string
	InputFile   string
	OutputPath  string
	Tags        map[string]TagInfo
}

func (ts *TemplateStruct) GenerateEventPkgFiles() error {

	var tmplFiles []string
	err := filepath.Walk("cmd/eventpkggen/tmpl/tmpl_files", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("err:%v\n", err)
		}
		if !info.IsDir() {
			tmplFiles = append(tmplFiles, path)
		}
		return nil
	})
	if err != nil {
		return err
	}

	t, err := template.ParseFiles(tmplFiles...)
	if err != nil {
		return fmt.Errorf("Failed to parse template files:%v\n", err)
	}

	for _, t := range t.Templates() {
		if _, err := os.Stat(ts.OutputPath); os.IsNotExist(err) {
			err := os.MkdirAll(ts.OutputPath, os.ModePerm)
			if err != nil {
				return fmt.Errorf("Failed to mkdir %v: %v\n", ts.OutputPath, err)
			}
		}
		var outputFile string
		switch {
		case strings.Contains(t.Name(), ".mod"):
			outputFile = path.Join(ts.OutputPath, strings.TrimSuffix(t.Name(), ".tmpl"))
		case t.Name() == "event_info.tmpl":
			eventsDir := path.Join(ts.OutputPath, "events")
			if _, err := os.Stat(eventsDir); os.IsNotExist(err) {
				err := os.Mkdir(eventsDir, os.ModePerm)
				if err != nil {
					return fmt.Errorf("Failed to generate dir %v: %v\n", eventsDir, err)
				}
			}
			outputFile = path.Join(eventsDir, strings.Join([]string{strings.TrimSuffix(t.Name(), ".tmpl"), ".go"}, ""))
		default:
			outputFile = path.Join(ts.OutputPath, strings.Join([]string{strings.TrimSuffix(t.Name(), ".tmpl"), ".go"}, ""))
		}

		f, err := os.Create(outputFile)
		if err != nil {
			return fmt.Errorf("error creating %v: %v\n", outputFile, err)
		}
		err = t.ExecuteTemplate(f, t.Name(), &ts)
		if err != nil {
			return fmt.Errorf("failed to execute template:%v\n", err)
		}
	}
	return nil
}
