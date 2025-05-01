build: tailwindcss gogenerate

gogenerate:
	go generate ./...

tailwindcss:
	npx tailwindcss -i templates/tailwind.css -o static/assets/css/styles.css --watch 

install-deps:
	npm install -D -g tailwindcss


dev: 
	templ generate --watch --cmd="go run main.go"
