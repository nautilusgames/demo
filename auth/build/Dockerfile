#build stage
FROM golang:alpine AS builder
WORKDIR /go/src/github.com/nautilusgames/demo
RUN apk add --no-cache git
ADD . /go/src/github.com/nautilusgames/demo
RUN go mod vendor
RUN go build -o ./auth/cmd/server ./auth/cmd/main.go

#final stage
FROM alpine:latest
ARG env=dev
RUN apk --no-cache add ca-certificates
RUN apk add --no-cache bash
COPY --from=builder /go/src/github.com/nautilusgames/demo/auth/cmd/server /app/server
COPY --from=builder /go/src/github.com/nautilusgames/demo/auth/configs/config.yaml /app/config.yaml

ADD https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh ./wait-for-it.sh
RUN ["chmod", "+x", "./wait-for-it.sh"]

EXPOSE 8080

CMD ["./app/server", "-c", "/app/config.yaml"]