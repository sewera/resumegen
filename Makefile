GEN_TEX=generated.tex
OUT_BIN=resumegen
GEN_DIR=./gen
OUT_PDF=$(GEN_DIR)/generated.pdf

LATEXMK_OPTIONS=-output-directory=$(GEN_DIR) -bibtex -pdf -pdflatex=pdflatex

.PHONY: build render clean

build:
	go build -o $(OUT_BIN) github.com/sewera/resumegen

render: $(OUT_PDF)

$(OUT_PDF): $(GEN_TEX)
	latexmk \
		$(LATEXMK_OPTIONS) \
		-f \
		-interaction=nonstopmode \
		$(GEN_TEX)
	@echo "> $@ rendered"

clean:
	rm -rf $(GEN_DIR)
	rm -f $(OUT_BIN)
	rm -f $(GEN_TEX)
	@echo "> $(GEN_DIR), $(OUT_BIN), and $(GEN_TEX) cleaned"
