FROM golang:1.18 AS builder
WORKDIR /go/src/github.com/vye/sleeper-go/
COPY . /go/src/github.com/vye/sleeper-go/
RUN go install
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/vye/sleeper-go/app ./
CMD ["./app"]  
