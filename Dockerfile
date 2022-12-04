FROM golang:alpine

RUN apk update && apk add --no-cache git && apk add --update make 

WORKDIR /usr/src/app

COPY . .

RUN go mod tidy

RUN go build -o ./app ./app/main.go

EXPOSE 8080

CMD ["./app/main"]