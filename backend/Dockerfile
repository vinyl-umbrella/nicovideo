# build
FROM golang:1.22.4 as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -o main .

# execute
FROM gcr.io/distroless/static-debian11:latest

ENV GIN_MODE=release
ENV TZ=Asia/Tokyo

WORKDIR /app
COPY --from=builder --chown=nonroot:nonroot /app/main .

USER nonroot:nonroot

CMD ["./main"]
