FROM alpine:latest

WORKDIR /app

COPY bin/main .

EXPOSE 3001

CMD ["./main"]