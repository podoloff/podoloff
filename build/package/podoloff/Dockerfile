FROM golang:alpine as builder

RUN apk add --no-cache ca-certificates git

# Download and install the latest release of dep
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

WORKDIR $GOPATH/src/github.com/podoloff/podoloff

COPY . ./

RUN dep ensure --vendor-only

RUN go build -o podoloff .

FROM alpine

RUN mkdir app

COPY --from=builder /go/src/github.com/podoloff/podoloff/podoloff /app/

WORKDIR /app

EXPOSE 8080

CMD ["./podoloff", "start"]