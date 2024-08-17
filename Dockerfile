#build stage
FROM golang:alpine AS builder
WORKDIR /movie_app
COPY . .
RUN go build -o movie_app

#final stage
FROM alpine:latest
COPY --from=builder /movie_app/movie_app .
ENTRYPOINT ["./movie_app"]
LABEL Name=movies Version=0.0.1
EXPOSE 8080