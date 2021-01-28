FROM golang:1.14

WORKDIR /app

RUN apt-get update && apt-get install git

RUN go get github.com/appleboy/gin-jwt

RUN go get github.com/gin-contrib/cors

RUN go get github.com/joho/godotenv

RUN go get gorm.io/gorm

RUN go get gorm.io/driver/postgres

RUN go get github.com/githubnemo/CompileDaemon

RUN go get github.com/confluentinc/confluent-kafka-go/kafka

# ENV CGO_ENABLED=0

ENTRYPOINT CompileDaemon -build="go build" -command="./app" -include=*.*

