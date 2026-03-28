FROM golang:1.26.1-alpine AS kobomatic-builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-s -w" \
    -trimpath \
    -o /kobomatic

# ------

FROM alpine:latest

COPY --from=kobomatic-builder /kobomatic /usr/local/bin/kobomatic

ENV SERVER_ADDRESS=""
ENV LIBRARY_FOLDER="/books"
ENV KOBOMATIC_FOLDER="/kobomatic"

EXPOSE 8084

CMD ["kobomatic"]