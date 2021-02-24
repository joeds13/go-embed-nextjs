run: start-client
	go run main.go -dev

build: export-client build-binary

build-binary:
	CGO_ENABLED=0 go build -ldflags "-w" -a -o go-embed-nextjs .

start-client:
	cd client; \
	npm run dev &

export-client:
	cd client; \
	rm -rf out/ && \
	npm run build && \
	npm run export && \
	find out/ -type d -exec touch {}/embed.txt \;

docker-build:
	docker build -t go-embed-nextjs:latest .

docker-run: docker-build
	docker run --rm -it -p 8080:8080 go-embed-nextjs:latest
