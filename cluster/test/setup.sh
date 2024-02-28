#!/usr/bin/env bash
echo "setup.sh"

SCRIPT_DIR=$( cd -- $( dirname -- "${BASH_SOURCE[0]}" ) &> /dev/null && pwd )

if [[ -n "${UPTEST_CLOUD_CREDENTIALS:-}" ]]; then
  echo "Creating cloud credential secret..."
  ${KUBECTL} -n upbound-system create secret generic provider-secret --from-literal=credentials="${UPTEST_CLOUD_CREDENTIALS}" --dry-run=client -o yaml | ${KUBECTL} apply -f -

  echo "Creating a default provider config..."
  cat <<EOF | ${KUBECTL} apply -f -
apiVersion: datadog.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: provider-secret
      namespace: upbound-system
      key: credentials
EOF
fi

${KUBECTL} wait provider.pkg --all --for condition=Healthy --timeout 5m
${KUBECTL} -n upbound-system wait --for=condition=Available deployment --all --timeout=5m
