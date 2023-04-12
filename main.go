package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

func main() {
	yamlFile := inputYAMLFile(inputFileName())
	resume := FromYAML(yamlFile)
	executeTemplate(resume)
	generateResume()
	renameResume(outputFileName())
}

func inputFileName() string {
	if len(os.Args) < 2 {
		return DefaultInputYAML
	}
	return os.Args[1]
}

func executeTemplate(resume Resume) {
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

func generateResume() {
	latexmkOptions := []string{
		fmt.Sprintf("-output-directory=%s", GeneratedDir),
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

func renameResume(target string) {
	err := os.Mkdir(OutputDir, 0755)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}

	err = os.Rename(GeneratedPDF, OutputDir+target)
	if err != nil {
		panic(err)
	}
}

func outputFileName() string {
	inputFile := inputFileName()
	baseFileName := filepath.Base(inputFile)
	outputFilename, ok := strings.CutSuffix(baseFileName, ".yaml")
	if !ok {
		outputFilename, _ = strings.CutSuffix(baseFileName, ".yml")
	}
	return outputFilename + ".pdf"
}
