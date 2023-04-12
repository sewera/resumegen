package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

func main() {
	yamlFile := inputYAMLFile(inputFileName())
	resume := FromYAML(yamlFile)
	executeTemplate(resume)
	generateResume()
}

func executeTemplate(resume Resume) {
	err := resumeTemplate().Execute(outputLatexFile(), resume)
	if err != nil {
		panic(err)
	}
}

func generateResume() {
	outputDir := "dist"
	latexmkOptions := []string{
		fmt.Sprintf("-output-directory=%s", outputDir),
		"-bibtex",
		"-pdf",
		"-pdflatex=pdflatex",
		"-f",
		"-interaction=nonstopmode",
		GeneratedLatex,
	}
	cmd := exec.Command("latexmk", latexmkOptions...)
	err := cmd.Run()
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
