FROM golang:alpine as builder

RUN mkdir /server

ADD . /server

WORKDIR /server

RUN go build -o podoloff .

FROM alpine

RUN adduser -S -D -H -h /app appuser

USER appuser

COPY --from=builder /server/podoloff /app/

WORKDIR /app

CMD ["./podoloff"]