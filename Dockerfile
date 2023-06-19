FROM golang:latest as builder

LABEL author="Yiwei Ge"

# Set the Current Working Directory inside the container
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo cmd/server/main.go


######## Start a new stage  #######

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .


EXPOSE 8081

CMD ["./main"]
