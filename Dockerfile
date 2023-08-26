FROM golang:1.20-alpine as builder

LABEL maintainer = "Matheus Carmo (a.k.a Carmel) <mematheuslc@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN ls

RUN GOOS=linux go build -o /app/api ./cmd/api

# Final image
FROM alpine:latest as runner
RUN apk --no-cache add ca-certificates

COPY --from=builder /app/api /api
EXPOSE 80

CMD /api