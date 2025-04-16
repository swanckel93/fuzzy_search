FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
RUN go install github.com/swaggo/swag/cmd/swag@latest && \
    export PATH=$PATH:/go/bin && \
    swag init
    
RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
