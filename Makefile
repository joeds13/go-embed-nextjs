run: start-client
	go run main.go -dev

build: export-client build-binary

build-binary:
	go generate
	CGO_ENABLED=0 go build -ldflags "-w" -a -o go-embed-nextjs .

start-client:
	cd client; \
	npm run dev &

export-client:
	cd client; \
	rm -rf out/ && \
	npm run build && \
	npm run export

docker-build:
	docker build -t go-embed-nextjs:latest .

docker-run: docker-build
	docker run --rm -it -p 8080:8080 go-embed-nextjs:latest

ensure-client-file-count-matches-go-embed-file-list:
	@echo "Number of exported Next.js files: "
	@find client/out -type f | wc -l
	@echo "Number of files for Go to embed: "
	@go list -json | jq -r .EmbedFiles[] | wc -l
