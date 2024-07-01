#build stage
FROM registry.nautilusgames.tech/marketplace/cicd/golang:1.22.1-builder-1.0 AS builder

#final stage
FROM alpine:3.15
RUN apk --no-cache add ca-certificates bash
ARG env=dev
ARG service
COPY --from=builder /tools/grpc_health_probe /grpc_health_probe
COPY --from=builder /tools/wait-for-it.sh ./wait-for-it.sh
COPY --from=builder /tools/envoy-preflight  /envoy-preflight
COPY out/cmd/main /app/server
COPY configs /app/
EXPOSE 8080

#For custom CMD scripts, please create scripts folder in $service folder, and update helm value
CMD ["/app/server", "-c", "/app/config.yaml"]
