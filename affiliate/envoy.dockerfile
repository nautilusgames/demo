FROM envoyproxy/envoy:v1.18.4 as builder
ARG service
#Backend dir context
WORKDIR /app
ADD $service /app/
ADD ./wait-for-it.sh /app/wait-for-it.sh
# Create if not exist
RUN mkdir -p /app/scripts && test -f /app/envoy.yaml || touch /app/envoy.yaml
FROM envoyproxy/envoy:v1.18.4
WORKDIR /app
# TODO remove on production
# net-tools dnsutils tcpdump vim traceroute : debuging tools
RUN apt-get update && apt-get install -y gettext net-tools dnsutils tcpdump vim traceroute \
  && rm -rf /var/lib/apt/lists/*
# For default
# Envoy sidecar helper
COPY ./envoy-preflight  /envoy-preflight
COPY --from=builder /app/wait-for-it.sh ./wait-for-it.sh
COPY --from=builder /app/envoy.yaml /etc/envoy.yaml
COPY --from=builder /app/descriptors /etc/descriptors
COPY --from=builder /app/scripts /app/scripts
RUN chmod +x -R /app/
EXPOSE 8080 8081

#For custom CMD scripts, please create scripts folder in $service folder, and update helm value