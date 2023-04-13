package main

import "os"

const ResumeTemplate = "resume.tex.gotpl"
const GeneratedLatex = "generated.tex"
const GeneratedDir = "gen/"
const GeneratedPDF = GeneratedDir + "generated.pdf"

func outputLatexFile() *os.File {
	file, err := writableCleanFile(GeneratedLatex)
	if err != nil {
		panic(err)
	}
	return file
}

func writableCleanFile(name string) (*os.File, error) {
	return os.OpenFile(name, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
}
