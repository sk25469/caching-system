FROM golang:1.19-alpine
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
EXPOSE 3000
RUN go build -o /app
CMD ["./main"]