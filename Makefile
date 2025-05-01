build: tailwindcss gogenerate

gogenerate:
	go generate ./...

tailwindcss:
	tailwindcss -i ./templates/tailwind.css -o ./static/assets/css/styles.css


install-deps:
	npm install -D -g tailwindcss


dev: 
	templ generate --watch --cmd="go run main.go"