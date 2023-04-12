package main

import (
	"os"
	"strings"
	"text/template"
)

func main() {
	yamlFile := inputYAMLFile(inputFileName())
	resume := FromYAML(yamlFile)
	err := resumeTemplate().Execute(outputLatexFile(), resume)

	if err != nil {
		panic(err)
	}
}

func inputFileName() string {
	if len(os.Args) < 2 {
		return DefaultInputYAML
	}
	return os.Args[1]
}

func resumeTemplate() *template.Template {
	return template.Must(template.
		New(ResumeTemplate).
		Delims("(((", ")))").
		Funcs(template.FuncMap{"join": strings.Join}).
		ParseFiles(ResumeTemplate))
}
