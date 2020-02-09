FROM golang:1.13.7 AS Test

ENV CGO_ENABLED=0
WORKDIR /go/src/app

RUN go get golang.org/x/tools/cmd/cover &&\
    go get github.com/mattn/goveralls

COPY . /go/src/app
RUN go test ./... -v -covermode=count -coverprofile=coverage.out
RUN go install

FROM alpine:3.9 as Deploy

EXPOSE 1323

COPY --from=Test /go/bin/mediapedia-api /app

ENTRYPOINT ["./app"]