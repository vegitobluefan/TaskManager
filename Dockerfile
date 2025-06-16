FROM golang:1.24.1 AS final

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server ./cmd

EXPOSE 8080

ENV DB_HOST=db
ENV DB_PORT=5432
ENV DB_USER=taskuser
ENV DB_PASSWORD=taskpass
ENV DB_NAME=taskdb

CMD ["./server"]
