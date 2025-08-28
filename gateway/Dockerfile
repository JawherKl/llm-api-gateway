FROM golang:1.24-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /bin/gateway ./cmd/server
FROM alpine:3.18
RUN apk add --no-cache ca-certificates
COPY --from=build /bin/gateway /bin/gateway
EXPOSE 3020
ENTRYPOINT ["/bin/gateway"]
