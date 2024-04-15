#Dockerfile

FROM golang:1.22.2-alpine
WORKDIR /app
COPY . .
RUN go build -o url_shortener .
EXPOSE 8080
CMD ["./url_shortener"]
