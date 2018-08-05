FROM golang:1.10 as builder

WORKDIR /go/src/cron
COPY . .

RUN go get -d -v ./...
RUN cd cmd && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo

FROM scratch
COPY --from=builder /go/src/cron/cmd/cmd ./
ENTRYPOINT ["./cmd"]