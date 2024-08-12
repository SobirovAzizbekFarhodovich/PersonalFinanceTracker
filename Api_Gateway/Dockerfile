FROM golang:1.22.3 AS builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/myapp .
COPY --from=builder /app/api/model.conf ./api/
COPY --from=builder /app/api/policy.csv ./api/

COPY .env .
EXPOSE 8082
CMD ["./myapp"]
