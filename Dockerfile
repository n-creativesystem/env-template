FROM golang:1.16-alpine as build

WORKDIR /src
COPY go.mod go.sum *.go ./
COPY docker-entrypoint.sh .

RUN go mod download \
    && go build -o cli \
    && chmod +x docker-entrypoint.sh \
    && chmod +x cli

FROM alpine

COPY --from=build /src/cli  /usr/bin/cli

COPY --from=build /src/docker-entrypoint.sh /usr/bin/docker-entrypoint.sh

ENTRYPOINT [ "docker-entrypoint.sh" ]