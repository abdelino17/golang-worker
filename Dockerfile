# Build stage
FROM golang:1.23-bookworm AS builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -v -o worker


# Run stage
FROM gcr.io/distroless/base-debian12

COPY --from=builder /app/worker /app/worker

CMD ["/app/worker"]