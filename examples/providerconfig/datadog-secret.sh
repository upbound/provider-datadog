cat <<EOF|kubectl apply -f -
apiVersion: v1
kind: Secret
metadata:
  name: datadog-creds
  namespace: upbound-system
type: Opaque
stringData:
  credentials: |
    {
      "api_key": "${DATADOG_API_KEY}",
      "app_key": "${DATADOG_APP_KEY}",
      "api_url": "https://api.datadoghq.com/"
    }
EOF
