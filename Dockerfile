FROM golang:latest
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main cmd/main.go
EXPOSE 3000
ENTRYPOINT [ "./main" ]