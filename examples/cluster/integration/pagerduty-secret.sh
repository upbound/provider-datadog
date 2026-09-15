cat <<EOF|kubectl apply -f -
apiVersion: v1
kind: Secret
metadata:
  name: pagerduty-secret
  namespace: upbound-system
type: Opaque
stringData:
  credentials: |
    {
      api_token: "${PAGERDUTY_INTEGRATION_KEY}"
    }
EOF
