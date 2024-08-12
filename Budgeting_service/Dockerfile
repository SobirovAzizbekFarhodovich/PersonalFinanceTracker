# Stage 1: Build the application
FROM golang:1.22.3-alpine AS build

WORKDIR /app

COPY . .

RUN go build -o budgeting-service

# Stage 2: Run the application
FROM alpine:latest

WORKDIR /root/

COPY --from=build /app/budgeting-service .

# Uncomment if .env file exists
COPY .env .

EXPOSE 50055

CMD ["./budgeting-service"]
