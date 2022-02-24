FROM golang:1.17.7-alpine AS build
RUN apk add --no-cache --update git

WORKDIR /go/src/github.com/claytonsilva/rabbitmq-go-service
# cache dependencies
COPY go.mod go.sum  ./
RUN go mod download
COPY . .
RUN go build ./cmd/main   && \
    chmod +x ./main

FROM alpine:latest
RUN apk add --no-cache ca-certificates

USER nobody
COPY --from=build /go/src/github.com/claytonsilva/rabbitmq-go-service/main /bin/rabbitmq-go-service
ENTRYPOINT ["/bin/rabbitmq-go-service"]
