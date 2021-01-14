FROM node:12.18.3-alpine3.12 AS JS_BUILD
COPY frontend /frontend
WORKDIR frontend
RUN npm install && npm run build

FROM golang:1.15.1-alpine3.12 AS GO_BUILD
RUN apk update && apk add build-base
COPY server /server
WORKDIR /server
RUN go build -o /go/bin/server

FROM alpine:3.12.0
COPY --from=JS_BUILD /webapp/build* ./frontend/
COPY --from=GO_BUILD /go/bin/server ./
CMD ./server
