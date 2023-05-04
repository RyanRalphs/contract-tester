FROM golang:1.18-bullseye


WORKDIR /app


COPY . .

WORKDIR /app/backend/cmd


RUN go build -o main .


EXPOSE 8080


CMD ["./main"]