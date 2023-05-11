FROM golang:1.20

WORKDIR /gin-exercise

COPY . ./
RUN go mod download
RUN go build -o out/http-server ./http_server

CMD ["./out/http-server"]