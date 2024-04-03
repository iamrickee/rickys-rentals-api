FROM golang:1.21.5 as base

FROM base as dev

ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /app
RUN mkdir "/build"
COPY . .
WORKDIR /app/src
RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon
RUN go get github.com/alexedwards/argon2id
RUN go get github.com/golang-migrate/migrate/v4/database/mysql
RUN go get github.com/joho/godotenv
RUN go get github.com/labstack/echo/v4
RUN go get github.com/labstack/echo/middleware
WORKDIR /app

ENTRYPOINT CompileDaemon -build="go build -o /build/app" -command="/build/app"

EXPOSE 8080