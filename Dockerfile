FROM golang:1.11.1 AS builder

COPY . /go/src/github.com/teran/microgpio

WORKDIR /go/src/github.com/teran/microgpio
RUN make dependencies build-linux-armv7


FROM scratch

ARG VCS_REF

LABEL org.label-schema.vcs-ref=$VCS_REF \
      org.label-schema.vcs-url="https://github.com/teran/microgpio"

COPY --from=builder /go/src/github.com/teran/microgpio/bin/microgpio-linux-armv7 /microgpio

ENTRYPOINT ["/microgpio"]
