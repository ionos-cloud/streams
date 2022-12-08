FROM alpine:latest as builder

RUN apk --update add ca-certificates

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY main-linux-amd64 /main

ENTRYPOINT ["/main"]
