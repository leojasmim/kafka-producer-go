# build stage
FROM golang:1.19.8 AS compiler
WORKDIR /kafka-producer-go
ADD go.mod go.sum /kafka-producer-go/
RUN GO111MODULE=on go mod download
ADD . .
RUN CGO_ENABLED=0 GOARCH=amd64 go build -o /kafka-producer-go/bin/application /kafka-producer-go/cmd/main.go

FROM alpine:3.10 AS production
ENV APP_API_PORT 3000
WORKDIR /kafka-producer-go
COPY --from=compiler /kafka-producer-go/bin/application /kafka-producer-go/application
RUN chmod +x /kafka-producer-go/application
RUN apk add --update --no-cache ca-certificates
ENTRYPOINT /kafka-producer-go/application
EXPOSE $APP_API_PORT
