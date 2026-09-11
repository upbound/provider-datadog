#!/usr/bin/env bash
set -aeuo pipefail

echo "Running setup.sh"

echo "Waiting until provider is healthy..."
${KUBECTL} wait provider.pkg --all --for condition=Healthy --timeout 5m

echo "Waiting for all pods to come online..."
${KUBECTL} -n crossplane-system wait --for=condition=Available deployment --all --timeout=5m

if [[ -n "${UPTEST_CLOUD_CREDENTIALS:-}" ]]; then
  echo "Creating cloud credential secret..."
  ${KUBECTL} -n crossplane-system create secret generic provider-secret --from-literal=credentials="${UPTEST_CLOUD_CREDENTIALS}" --dry-run=client -o yaml | ${KUBECTL} apply -f -

  echo "Creating a default provider config..."
  cat <<YAML | ${KUBECTL} apply -f -
apiVersion: datadog.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: provider-secret
      namespace: crossplane-system
      key: credentials
YAML

  echo "Creating a default cluster provider config (v2-style)..."
  cat <<YAML | ${KUBECTL} apply -f -
apiVersion: datadog.m.upbound.io/v1beta1
kind: ClusterProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: provider-secret
      namespace: crossplane-system
      key: credentials
YAML
fi

${KUBECTL} wait provider.pkg --all --for condition=Healthy --timeout 5m
${KUBECTL} -n crossplane-system wait --for=condition=Available deployment --all --timeout=5m
