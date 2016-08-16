FROM golang:1.6.3-alpine
MAINTAINER Jordan Earls <jordanearls@gmail.com>
RUN apk --no-cache add git
RUN go get -u github.com/kataras/iris/iris
RUN mkdir -p /go/src/app
WORKDIR /go/src/app

# this will ideally be built by the ONBUILD below ;)
CMD ["go-wrapper", "run"]

ONBUILD COPY . /go/src/app
ONBUILD RUN go-wrapper download
ONBUILD RUN go-wrapper install

EXPOSE 8080
ENTRYPOINT  ["go", "run", "app.go"]
