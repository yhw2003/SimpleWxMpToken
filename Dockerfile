FROM hub.atomgit.com/library/golang:1-alpine AS builder
COPY . /app
WORKDIR /app
RUN go build -o bin
FROM hub.atomgit.com/amd64/alpine:3.15.10
COPY --from=builder /app/bin /app/bin
COPY ./.env /app/.env
ENV LANG C.UTF-8
WORKDIR /app
CMD ["./bin"]