.PHONY: build-css
build-css:
	npx tailwindcss -i src/css/app.postcss -o public/styles.css

.PHONY: gen-views
gen-views:
	templ generate internal/view

.PHONY: install-tool
install-tool:
	go install github.com/a-h/templ/cmd/templ@latest

.phony: tidy
tidy:
	go mod tidy
	go mod vendor