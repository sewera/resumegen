package main

import "os"

const ResumeTemplate = "resume.tex.gotpl"
const DefaultInputYAML = "example.yaml"
const GeneratedLatex = "generated.tex"

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
	var fileName string
	if name == "" {
		fileName = DefaultInputYAML
	} else {
		fileName = name
	}

	file, err := readableFile(fileName)
	if err != nil {
		panic(err)
	}
	return file
}

func readableFile(name string) (*os.File, error) {
	return os.Open(name)
}
