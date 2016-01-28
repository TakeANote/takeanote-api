FROM golang:1.5

RUN mkdir -p /go/src/github.com/takeanote/takeanote-api

WORKDIR /go/src/github.com/takeanote/takeanote-api

COPY . /go/src/github.com/takeanote/takeanote-api

RUN go-wrapper download
RUN go-wrapper install

EXPOSE 80 443

CMD ["go-wrapper", "run"]
