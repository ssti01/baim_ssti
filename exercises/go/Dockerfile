FROM golang:1.23-alpine3.21

WORKDIR /app

COPY main.go .

RUN go build main.go

FROM alpine:3.21

WORKDIR /app

COPY --from=0 /app/main .

COPY index.html .

EXPOSE 5555

CMD ["./main"]