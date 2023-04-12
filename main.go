package main

import (
	"os"
	"strings"
	"text/template"
)

func main() {
	err := resumeTemplate().Execute(outputLatexFile(), Example)

	if err != nil {
		panic(err)
	}
}

func resumeTemplate() *template.Template {
	return template.Must(template.
		New("resume.tex.gotpl").
		Delims("(((", ")))").
		Funcs(template.FuncMap{"join": strings.Join}).
		ParseFiles("resume.tex.gotpl"))
}

func outputLatexFile() *os.File {
	file, err := writableCleanFile("generated.tex")
	if err != nil {
		panic(err)
	}
	return file
}

func writableCleanFile(name string) (*os.File, error) {
	return os.OpenFile(name, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
}
