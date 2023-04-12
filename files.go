package main

import "os"

const ResumeTemplate = "resume.tex.gotpl"
const DefaultInputYAML = "example.yaml"
const GeneratedLatex = "generated.tex"
const GeneratedDir = "gen/"
const GeneratedPDF = GeneratedDir + "generated.pdf"
const OutputDir = "dist/"

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

func inputYAMLFile(name string) *os.File {
	file, err := readableFile(name)
	if err != nil {
		panic(err)
	}
	return file
}

func readableFile(name string) (*os.File, error) {
	return os.Open(name)
}
