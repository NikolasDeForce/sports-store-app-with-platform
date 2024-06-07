FROM golang:latest

COPY . /app

EXPOSE 5500
WORKDIR /app/sportsstore
ENTRYPOINT [ "go", "run", "main.go" ]