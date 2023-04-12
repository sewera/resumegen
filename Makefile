GEN_TEX=generated.tex
OUT_BIN=resumegen
OUT_DIR=./dist
OUT_PDF=$(OUT_DIR)/generated.pdf

LATEXMK_OPTIONS=-output-directory=$(OUT_DIR) -bibtex -pdf -pdflatex=pdflatex

.PHONY: clean

$(OUT_PDF): $(GEN_TEX)
	@latexmk \
		$(LATEXMK_OPTIONS) \
		-synctex=1 \
		-f \
		-interaction=nonstopmode \
		$(GEN_TEX) > /dev/null
	@echo "> $@ rendered"

clean:
	@rm -rf $(OUT_DIR)
	@rm -f $(OUT_BIN)
	@rm -f $(GEN_TEX)
	@echo "> $(OUT_DIR), $(OUT_BIN), and $(GEN_TEX) cleaned"

build:
	@go build -o $(OUT_BIN) github.com/sewera/resume-generator-latex