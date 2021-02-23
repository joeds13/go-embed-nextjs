FROM node:15 AS build-ui
WORKDIR /build
COPY . /build
RUN make export-ui

FROM golang:1.16 AS build-binary
WORKDIR /build
COPY --from=build-ui /build /build
RUN make build-binary

FROM scratch
COPY --from=build-binary /build/go-embed-nextjs /
CMD ["/go-embed-nextjs"]
