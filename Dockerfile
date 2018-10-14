FROM golang:1.11.1 AS builder

COPY . /go/src/github.com/teran/microgpio

WORKDIR /go/src/github.com/teran/microgpio
RUN make dependencies build-linux-amd64 build-linux-armv7


FROM alpine:3.8

ARG VCS_REF

LABEL org.label-schema.vcs-ref=$VCS_REF \
      org.label-schema.vcs-url="https://github.com/teran/microgpio"

RUN apk add --update --no-cache \
  ca-certificates=20171114-r3 && \
  rm -vf /var/cache/apk/*

COPY --from=builder /go/src/github.com/teran/microgpio/bin/microgpio-linux-amd64 /microgpio

ENTRYPOINT ["/microgpio"]
