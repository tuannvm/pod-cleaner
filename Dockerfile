FROM golang:1.11-alpine AS builder


RUN apk add --no-cache git curl
COPY . /go/src/pod-cleaner
WORKDIR /go/src/pod-cleaner

# https://github.com/golang/dep/blob/master/docs/FAQ.md#how-do-i-use-dep-with-docker
RUN set -ex \
    && go get \
    && go build -v -o "/pod-cleaner"

# -- pod-cleaner Image --
FROM alpine:3.7
RUN set -ex \
    && apk add --no-cache bash ca-certificates git

COPY --from=builder /pod-cleaner /bin/pod-cleaner
ENTRYPOINT [ "/bin/pod-cleaner" ]
