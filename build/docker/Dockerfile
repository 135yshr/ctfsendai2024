FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o main ./main.go


FROM scratch

EXPOSE 8080

WORKDIR /app

COPY --from=builder /app/main .

COPY --from=builder /app/configs ./configs
COPY --from=builder /app/web ./web

ENTRYPOINT ["./main"]
