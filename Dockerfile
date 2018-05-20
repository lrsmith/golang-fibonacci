FROM golang:onbuild

MAINTAINER Len Smith "lrsmith@umich.edu"
EXPOSE 8443/tcp

WORKDIR /go/src/app
RUN go build 
ENTRYPOINT /go/src/app/app

HEALTHCHECK --interval=5m --timeout=3s \
  CMD curl --insecure https://localhost:8443/status || exit 1

