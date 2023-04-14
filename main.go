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
	resume := FromYAML(os.Stdin)
	executeTemplate(resume)
	generateResume()
	if name, shouldRename := outputFileName(); shouldRename {
		renameResume(name)
	}
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
		ParseFS(resumeTemplateFS, ResumeTemplate))
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
	dir := filepath.Dir(target)
	err := os.Mkdir(dir, 0755)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}

	err = os.Rename(GeneratedPDF, target)
	if err != nil {
		panic(err)
	}
}

func outputFileName() (string, bool) {
	if len(os.Args) < 2 {
		return "", false
	}
	return os.Args[1], true
}
