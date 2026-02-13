# Stage 1: Build Hugo site
FROM docker.io/hugomods/hugo:exts AS hugo-builder
WORKDIR /src
COPY . .
RUN hugo --minify

# Stage 2: Build Go server
FROM docker.io/library/golang:1.24-alpine AS go-builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY cmd/server/ cmd/server/
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /server ./cmd/server

# Stage 3: Final minimal image
FROM gcr.io/distroless/static:nonroot
COPY --from=go-builder /server /server
COPY --from=hugo-builder /src/public /site
EXPOSE 8080
ENTRYPOINT ["/server"]
