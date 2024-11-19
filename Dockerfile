FROM golang:1.23 as builder
ADD . /go/src/github.com/silenceper/go-mock-demo/
RUN cd /go/src/github.com/silenceper/go-mock-demo/ \
  && go get -v \
  && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM centos:7
MAINTAINER silenceper <silenceper@gmail.com>
COPY --from=builder /go/src/github.com/silenceper/go-mock-demo/app /bin/app
ENTRYPOINT ["/bin/app"]
EXPOSE 80