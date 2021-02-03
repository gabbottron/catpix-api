# This is based on my generic golang two-stage
# Docker build. Set variables in docker-compose.

### ------ BUILD STAGE -----------------------------
FROM golang:1.15 as build-stage
WORKDIR /wrk/catpix-api
COPY . .
RUN CGO_ENABLED=0 go build -o ./bin/catpix-api -a -installsuffix cgo ./src/catpix-api/main.go

### ------- RUN STAGE ------------------------------
FROM alpine:latest as run-stage
RUN apk --no-cache add ca-certificates
RUN apk --no-cache add postgresql-client
WORKDIR /app/
COPY --from=build-stage /wrk/catpix-api/bin/catpix-api .
