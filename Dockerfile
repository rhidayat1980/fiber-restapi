FROM golang:alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN GOOS=linux GOARCH=amd64 go build -o app . 

EXPOSE 3000

CMD [ "./app" ]