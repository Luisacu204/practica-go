FROM golang:1.21.5-alpine 


WORKDIR /app

COPY go.mod go.sum ./

COPY . .

COPY *.go ./
  
COPY prueba.txt .

# RUN mkdir bin

RUN go mod download

# Expose port 4000 to the outside world
EXPOSE 8080

RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]
RUN ["go", "install", "github.com/githubnemo/CompileDaemon"]


ENTRYPOINT CompileDaemon -polling -log-prefix=false -build="go build -o ./bin/" -command="./bin/proyecto" -directory="./"