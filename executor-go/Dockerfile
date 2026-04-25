FROM golang:1.25

WORKDIR /app
COPY . .

RUN apt-get update && apt-get install -y \
    python3 \
    default-jdk \
    g++

RUN go mod tidy
RUN go build -o app

CMD ["./app"]