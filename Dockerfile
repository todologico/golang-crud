FROM golang:1.22-alpine

WORKDIR /app

RUN go install github.com/githubnemo/CompileDaemon@latest

COPY go.mod go.sum ./

RUN apk add --no-cache git

RUN go mod download && go mod verify

COPY *.go ./

RUN go build -o main . 

RUN chmod +x /app/main

EXPOSE 8080

#CMD ["CompileDaemon", "-build=go run main.go", "-directory=/app"]
#CMD ["CompileDaemon", "-build=go build -o /app", "-directory=/app"]

ENTRYPOINT ["CompileDaemon", "--build=go build -o main .", "--command=./main"]
