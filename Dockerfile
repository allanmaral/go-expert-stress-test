# Builder
FROM golang:latest as builder

WORKDIR /app

COPY go.mod go.sum* ./
RUN go mod download

COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o stress .

# Runner
FROM scratch

COPY --from=builder /app/stress /usr/bin/stress

ENTRYPOINT [ "/usr/bin/stress" ]