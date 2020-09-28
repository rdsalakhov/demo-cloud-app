FROM golang:alpine

RUN mkdir /app
WORKDIR /app
COPY . .
RUN go build -o ./main
ENTRYPOINT ["./main"]
