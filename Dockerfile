FROM golang:1.13.7 AS Test

ENV CGO_ENABLED=0
WORKDIR /go/src/app
COPY . /go/src/app

RUN go get golang.org/x/tools/cmd/cover &&\
    go get github.com/mattn/goveralls

RUN go test ./... -v -covermode=count -coverprofile=coverage.out
RUN go install

FROM alpine:3.9 as Deploy
COPY --from=Test /go/bin/mediapedia-api /app

EXPOSE 1323

ENTRYPOINT ["./app"]
