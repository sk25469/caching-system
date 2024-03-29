
# syntax=docker/dockerfile:1

FROM golang:1.19-alpine
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
EXPOSE 3000
RUN go build /app
CMD ["./main"]