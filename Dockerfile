FROM golang:1.17-alpine as builder

# building app
WORKDIR /go/src/app
ADD . /go/src/app
RUN go get -d -v
RUN go build -o /go/bin/app main.go

FROM alpine:3.16
COPY --from=builder /go/bin/app /
ENTRYPOINT ["/app"]
