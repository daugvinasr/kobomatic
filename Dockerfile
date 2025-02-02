FROM golang:1.23.2-alpine AS kobomatic-builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /kobomatic

# ------

FROM golang:1.23.2-alpine AS kepubify-builder
WORKDIR /app

RUN apk add --no-cache git && \
    git clone --depth 1 https://github.com/pgaskin/kepubify.git && \
    cd kepubify && \
    CGO_ENABLED=0 GOOS=linux go build -o /kepubify ./cmd/kepubify

# ------

FROM alpine:latest

COPY --from=kobomatic-builder /kobomatic /usr/local/bin/kobomatic
COPY --from=kepubify-builder /kepubify /usr/local/bin/kepubify

ENV SERVER_ADDRESS=""
ENV LIBRARY_FOLDER="/books"
ENV CKSS_FOLDER="/kobomatic"

EXPOSE 8084

CMD ["kobomatic"]