FROM golang:1.24-alpine 

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod /
RUN go mod download

COPY . .

RUN go build -o api-gateway .

EXPOSE 8080

CMD ["./api-gateway"]
