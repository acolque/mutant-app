FROM golang:1.15-alpine AS GO_BUILD
COPY . /server
WORKDIR /server/api
RUN go build -o /go/bin/server/api

FROM alpine:3.10
WORKDIR app
COPY --from=GO_BUILD /go/bin/server/ ./
EXPOSE 8080
CMD ./api