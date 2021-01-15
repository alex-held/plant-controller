FROM golang:1.15.1-alpine3.12 AS GO_BUILD
RUN apk update && apk add build-base
COPY . /build
WORKDIR /build
RUN go build -o /go/bin/server

FROM alpine:3.12.0
EXPOSE 8080
COPY --from=GO_BUILD /go/bin/server ./
CMD ./server
