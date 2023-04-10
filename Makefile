SRC_TEX=main.tex
OUT_DIR=./dist
OUT_PDF=$(OUT_DIR)/main.pdf

LATEXMK_OPTIONS=-output-directory=$(OUT_DIR) -bibtex -pdf -pdflatex=pdflatex

.PHONY: render clean

render: $(OUT_PDF)

$(OUT_PDF): $(SRC_TEX)
	@latexmk \
		$(LATEXMK_OPTIONS) \
		-synctex=1 \
		-f \
		-interaction=nonstopmode \
		$(SRC_TEX) > /dev/null
	@echo "> $@ rendered"

clean:
	@rm -rf $(OUT_DIR)
	@echo "> $(OUT_DIR) cleaned"
