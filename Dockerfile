FROM golang:alpine3.20

WORKDIR /app

RUN go install github.com/githubnemo/CompileDaemon@latest

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -o main .

RUN chmod +x /app/main

EXPOSE 8080

ENTRYPOINT ["CompileDaemon", "--build=go build -o main .", "--command=./main"]
