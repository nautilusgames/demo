#build stage
FROM registry.nautilusgames.tech/marketplace/cicd/golang:1.22.1-builder-1.0 AS builder
#Backend dir context
ARG BOT_USER
ARG BOT_PRIVATE_TOKEN
WORKDIR /go/app
ADD go.mod /go/app/go.mod
RUN echo "machine gitlab.nautilusgames.tech" >> ~/.netrc && \
    echo "login $BOT_USER" >> ~/.netrc && \
    echo "password $BOT_PRIVATE_TOKEN" >> ~/.netrc
RUN git config --global url."https://${BOT_USER}:${BOT_PRIVATE_TOKEN}@gitlab.nautilusgames.tech/".insteadOf https://gitlab.nautilusgames.tech/ &&\
    go env -w GOPRIVATE=gitlab.nautilusgames.tech && \
    go mod download all && cp go.sum /tmp/go.sum
ADD . /go/app
ARG service
RUN cp /tmp/go.sum ./go.sum && go build -mod=readonly  -o ./cmd/main ./cmd/main.go && mkdir -p ./scripts

#final stage
FROM alpine:3.15
ARG env=dev
ARG service
RUN apk --no-cache add ca-certificates bash
COPY --from=builder /go/app/cmd/main /app/server
COPY --from=builder /go/app/configs /app
COPY --from=builder /go/app/scripts /app/scripts

EXPOSE 8080

#For custom CMD scripts, please create scripts folder in $service folder, and update helm value
CMD ["./app/server", "-c", "/app/config.yaml"]
