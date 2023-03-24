package tmpl

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path"
	"strings"
	"text/template"
)

var (
	//go:embed tmpl_files/*
	tmpl embed.FS
)

type TagInfo struct {
	Tag            string
	TagName        string
	TagPackageName string
	TagImportName  string
}
type TemplateStruct struct {
	PackageName          string
	InputFile            string
	OutputPath           string
	LatestTagPackageName string
	Tags                 map[string]TagInfo
}

func (ts *TemplateStruct) GenerateEventPkgFiles() error {

	var tmplFiles []string
	err := fs.WalkDir(tmpl, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("err:%v\n", err)
		}
		if !d.IsDir() {
			tmplFiles = append(tmplFiles, path)
		}
		return nil
	})
	if err != nil {
		return err
	}

	t, err := template.ParseFS(tmpl, tmplFiles...)
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
