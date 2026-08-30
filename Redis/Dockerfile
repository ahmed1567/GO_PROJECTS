# Multi-stage build, same pattern as HttpServer/Dockerfile: compile in a
# full Go image, ship only the static binary. This project has an actual
# dependency (go-redis) hence go.sum exists — download it as its own layer
# before copying source, so editing a .go file doesn't bust the module cache.
FROM golang:1.26-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 go build -o /app-bin .

FROM alpine:3.20
COPY --from=build /app-bin /app-bin
ENTRYPOINT ["/app-bin"]
