FROM golang:1.6.3-alpine
MAINTAINER Jordan Earls <jordanearls@gmail.com>
RUN apk --no-cache add git
RUN go get -u github.com/kataras/iris/iris
RUN mkdir -p /go/src/app
WORKDIR /go/src/app

COPY . /go/src/app
RUN go-wrapper download
RUN go-wrapper install

EXPOSE 8080
ENTRYPOINT [ "go-wrapper", "run", "app.go" ]
