run: export-ui
	go run main.go

build: export-ui build-binary

build-binary:
	CGO_ENABLED=0 go build -ldflags "-w" -a -o go-embed-nextjs .

export-ui:
	cd client; \
	rm -rf out/ && \
	npm run build && \
	npm run export && \
	find out/ -type d -exec touch {}/embed.txt \;

docker-build:
	docker build -t go-embed-nextjs:latest .

docker-run: docker-build
	docker run --rm -it -p 8080:3000 go-embed-nextjs:latest
