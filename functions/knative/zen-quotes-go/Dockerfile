FROM golang:1.21-alpine AS build
COPY  . /app
WORKDIR /app
RUN go build -o go-quote-app

FROM alpine:latest
COPY --from=build /app/go-quote-app .
EXPOSE 8080
ENTRYPOINT [ "./go-quote-app" ]

RUN apk add --no-cache bash curl