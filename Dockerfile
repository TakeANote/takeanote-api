FROM golang:1.5-alpine

RUN apk add --update git && rm -rf /var/cache/apk/*

RUN mkdir -p $GOPATH/src/github.com/takeanote/takeanote-api

WORKDIR $GOPATH/src/github.com/takeanote/takeanote-api

COPY . $GOPATH/src/github.com/takeanote/takeanote-api

RUN go get -v ./... && \
    CGO_ENABLED=0 GOOS=linux go install -a -tags netgo -ldflags '-w' .

EXPOSE 80 443

CMD ["takeanote-api"]
