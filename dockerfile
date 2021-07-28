FROM golang:alpine AS builder

WORKDIR /app

COPY go.sum go.mod ./

RUN go mod download

COPY . ./

RUN go build -o main ./main.go

FROM golang:alpine

WORKDIR /app

COPY --from=builder /app/main /app/

EXPOSE 5000

CMD ["/app/main"]