🛠️ File Setup

1. Configuration Directory

Create vault/config folder. (This example doesn’t need extra config files, but it’s useful if switching to non-dev mode later.)

2. Persistent Data

Ensure vault/data exists — it’s where Vault stores data if restarted. Using file storage backend like this makes the setup more durable. ￼

✅ Start Vault
```bash
docker-compose up -d
```

✅ Use Vault CLI
```bash
export VAULT_ADDR=http://127.0.0.1:8200
export VAULT_TOKEN=devroot
vault status
```

✅ Check Vault UI
```bash
http://localhost:8200/ui/vault/ui/vault/ui/
```

✅ Store Gemini API Key & Ad Platform Secrets
```bash
vault kv put secret/adflowhq/gemini api_key=your-gemini-api-key
vault kv put secret/adflowhq/google-ads google_ads_key=your-google-ads-key
vault kv put secret/adflowhq/facebook-ads facebook_ads_key=your-facebook-ads-key
```

```bash
vault kv put secret/gemini api_key="YOUR_GEMINI_API_KEY"
```

```bash
vault kv put secret/google_ads client_id="YOUR_CLIENT_ID" client_secret="YOUR_CLIENT_SECRET" refresh_token="YOUR_REFRESH_TOKEN"
```

```bash
vault kv put secret/facebook_ads app_id="YOUR_APP_ID" app_secret="YOUR_APP_SECRET" access_token="YOUR_ACCESS_TOKEN"
```

```bash
vault kv put secret/callrail api_key="YOUR_CALLRAIL_API_KEY"
```
    








