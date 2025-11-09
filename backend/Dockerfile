FROM golang:1.22-alpine
WORKDIR /app
COPY . .
RUN apk add --no-cache build-base
RUN go mod tidy
RUN go build -o main .
EXPOSE 8080
CMD ["./main"]
