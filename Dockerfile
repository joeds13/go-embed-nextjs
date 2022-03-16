FROM node:17 AS build-client
ENV NEXT_TELEMETRY_DISABLED 1
WORKDIR /build
COPY Makefile /build/
COPY client/package.json /build/client/
RUN make install-client
COPY . /build
RUN make export-client

FROM golang:1.18 AS build-binary
WORKDIR /build
COPY --from=build-client /build /build
RUN make build-binary

FROM scratch
COPY --from=build-binary /build/go-embed-nextjs /
CMD ["/go-embed-nextjs"]
EXPOSE 8080
