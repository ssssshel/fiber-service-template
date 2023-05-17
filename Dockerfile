# BUILD STAGE

FROM golang:1.19.4-alpine AS builder

WORKDIR /app

COPY . /app

RUN go build -o main .

# FINAL STAGE

FROM alpine:3.14

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 3000-3001

USER nonroot:nonroot

CMD ["./main"]