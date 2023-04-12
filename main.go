package main

import (
	"strings"
	"text/template"
)

func main() {
	yamlFile := inputYAMLFile("")
	resume := FromYAML(yamlFile)
	err := resumeTemplate().Execute(outputLatexFile(), resume)

	if err != nil {
		panic(err)
	}
}

func resumeTemplate() *template.Template {
	return template.Must(template.
		New(ResumeTemplate).
		Delims("(((", ")))").
		Funcs(template.FuncMap{"join": strings.Join}).
		ParseFiles(ResumeTemplate))
}
