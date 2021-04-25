FROM golang:latest
WORKDIR /app

RUN go get -v github.com/google/uuid
RUN go get -v github.com/gorilla/websocket

COPY . .

RUN go build -o main .

ENV WORD=hackthepig
ENV FLAG="flag{4A4D_4ZH2gQyfumFR-ME97BxGHhSXNCC}"

EXPOSE 8080

CMD ["./main"]
