# Multi-stage build: compile in a full Go image, ship only the tiny
# static binary — a real advantage over Node images, which have to ship
# node_modules and the whole Node runtime alongside your code.
FROM golang:1.26-alpine AS build
WORKDIR /app
COPY go.mod ./
COPY cmd ./cmd
COPY internal ./internal
RUN CGO_ENABLED=0 go build -o /server ./cmd/server

FROM alpine:3.20
COPY --from=build /server /server
EXPOSE 8080
ENTRYPOINT ["/server"]
