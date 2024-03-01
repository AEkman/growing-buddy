FROM docker.io/golang:1.22.0-bullseye as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -o app ./cmd

FROM gcr.io/distroless/base-debian12:nonroot

WORKDIR /

COPY --from=builder /app/app .
EXPOSE 8333
ENTRYPOINT ["/app"]

