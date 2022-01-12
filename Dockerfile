FROM golang:1.17.1-alpine

WORKDIR /app

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY cmd ./cmd
COPY internal ./internal

RUN go build -o /norma-http

EXPOSE 8080

CMD [ "/norma-http" ]