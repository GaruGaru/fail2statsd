FROM golang

COPY ./ /app
WORKDIR /app

RUN mkdir /logs

VOLUME /logs

RUN go get github.com/hpcloud/tail/...

RUN go get github.com/cactus/go-statsd-client/statsd

CMD ["go","run","main.go"]