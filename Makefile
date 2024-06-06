run: build
	@./bin/app

build:
	npx tailwindcss -i views/css/app.css -o public/styles.css
	@templ generate view
	@go build -o bin/app .

tailwind:
	npx tailwindcss -i views/css/app.css -o public/styles.css --watch

templ:
	templ generate -watch -proxy=http://localhost:4000