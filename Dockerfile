FROM golang:1.20.3 AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /dist .

FROM gcr.io/distroless/base-debian11

WORKDIR /src

COPY --from=builder /dist /dist

CMD ["/dist"]
