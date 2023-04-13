# Resume Generator LaTeX
A tool for generating a classic style resume from metadata, templated in LaTeX.

## Prerequisites

- Make
- Go
- LaTeX tools (latexmk)

## Run

```sh
make build
cat <input_file.yaml> | ./resumegen <output_file.pdf>
```

## Combining inputs

You can have structured files for different kinds of resumes,
like:

```sh
cat base.yaml company1.yaml backend.yaml | ./resumegen your_name_company1_backend.pdf
```