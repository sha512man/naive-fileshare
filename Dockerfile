FROM golang:1.22-alpine AS build

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o naive-fileshare .

# ================
FROM alpine:latest

WORKDIR /app

COPY --from=build /app/naive-fileshare .

CMD ["/app/naive-fileshare"]
