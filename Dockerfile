FROM golang:alpine AS build

WORKDIR /build

COPY . .

RUN go mod download && \
    go build -o ./hello-world

FROM alpine

WORKDIR /app

COPY --from=build /build/hello-world ./hello-world

EXPOSE 8080

CMD ["./hello-world"]
