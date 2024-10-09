#!/bin/sh
set -e

echo "Generating envoy.yaml config file..."
cat /app/scripts/envoy.yaml | envsubst > /etc/envoy.yaml

echo "Starting Envoy..."
/usr/local/bin/envoy -c /etc/envoy.yaml
