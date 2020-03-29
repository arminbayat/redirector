FROM golang:1.13.1

WORKDIR /app/go-sample-app
COPY . .
RUN go build -o ./main .
EXPOSE 9091

CMD ["./main"]