FROM golang:1.12-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git

ADD . /tomochain
RUN cd /tomochain && make tao

FROM alpine:latest

WORKDIR /tomochain

COPY --from=builder /tomochain/build/bin/tao /usr/local/bin/tao

RUN chmod +x /usr/local/bin/tao

EXPOSE 8545
EXPOSE 20202

ENTRYPOINT ["/usr/local/bin/tao"]

CMD ["--help"]
