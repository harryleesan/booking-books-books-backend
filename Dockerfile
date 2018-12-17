FROM golang:1.11.3-stretch
WORKDIR /go/src/booking-books-books-backend
RUN apt-get update && apt-get install curl -y
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
COPY Gopkg.toml .
COPY Gopkg.lock .
RUN dep ensure -vendor-only
COPY cmd ./cmd
COPY pkg ./pkg
# RUN go build main.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/app/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root
COPY --from=0 /go/src/booking-books-books-backend/main .
CMD ["./main"]
