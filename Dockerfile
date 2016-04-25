FROM golang:1.6-alpine

RUN apk add --update git && rm -rf /var/cache/apk/*
RUN go get -u github.com/Masterminds/glide

WORKDIR $GOPATH/src/github.com/takeanote/takeanote-api
COPY . $GOPATH/src/github.com/takeanote/takeanote-api

RUN glide install
RUN CGO_ENABLED=0 GOOS=linux go build -o takeanote-api -a -tags netgo -ldflags '-w' ./cmd/takeanote-api

EXPOSE 80 443

CMD ["./takeanote-api"]
