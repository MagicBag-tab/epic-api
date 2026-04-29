FROM golang:1.22-alpine

RUN apk add --no-cache gcc musl-dev sqlite

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o epic-api .

EXPOSE 8088

CMD ["./epic-api"]