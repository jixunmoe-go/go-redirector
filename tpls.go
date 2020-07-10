package main

import (
	"bytes"
	"fmt"
	"text/template"
)

var tplImportURL *template.Template
var tplBaseWebURL *template.Template
var tplDirectoryURL *template.Template
var tplFileURL *template.Template
var tplSourceURL *template.Template

func initTemplates() {
	tplImportURL = template.Must(template.New("ImportURL").Parse(config.ImportURL))
	tplBaseWebURL = template.Must(template.New("BaseWebURL").Parse(config.BaseWebURL))
	tplDirectoryURL = template.Must(template.New("DirectoryURL").Parse(config.DirectoryURL))
	tplFileURL = template.Must(template.New("FileURL").Parse(config.FileURL))
	tplSourceURL = template.Must(template.New("SourceURL").Parse(config.SourceURL))
}

func textTemplateToString(t *template.Template, data interface{}) string {
	var buf bytes.Buffer
	err := t.Execute(&buf, data)
	if err != nil {
		fmt.Printf("Could not execute template: %s\n", err)
		return ""
	}
	return buf.String()
}
