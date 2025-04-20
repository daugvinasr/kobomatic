FROM golang:1.24.2-alpine AS kobomatic-builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /kobomatic

# ------

FROM alpine:latest

COPY --from=kobomatic-builder /kobomatic /usr/local/bin/kobomatic

ENV SERVER_ADDRESS=""
ENV LIBRARY_FOLDER="/books"
ENV KOBOMATIC_FOLDER="/kobomatic"

EXPOSE 8084

CMD ["kobomatic"]