#Dockerfile

FROM golang:1.22.2-alpine
WORKDIR /app
COPY . .
RUN go build -o URL_SHORTENER .
EXPOSE 8080
CMD ["./url-shortener"]
