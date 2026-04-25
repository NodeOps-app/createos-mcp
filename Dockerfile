FROM golang:1.26-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/createos-mcp .

FROM gcr.io/distroless/static:nonroot
COPY --from=build /out/createos-mcp /createos-mcp
EXPOSE 8080 9090
USER nonroot:nonroot
ENTRYPOINT ["/createos-mcp"]
