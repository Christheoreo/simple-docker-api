FROM golang:1-alpine
WORKDIR /app
COPY . .
EXPOSE 1323
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app .
CMD ./app