# Dev stage
FROM golang:1.21 AS dev

ENV APP_HOME /go/src/github.com/michelsazevedo/authz/
WORKDIR $APP_HOME

COPY go.mod ./

RUN go mod download && go mod verify

COPY . .

# Builder stage
FROM dev AS builder

ENV APP_HOME /go/src/github.com/michelsazevedo/authz/
WORKDIR $APP_HOME

RUN CGO_ENABLED=0 GOOS=linux go build -o authz .

# Production stage
FROM alpine:latest AS production

ENV APP_HOME /go/src/github.com/michelsazevedo/authz/

COPY --from=builder $APP_HOME .

EXPOSE 8080

CMD ["./authz"]
