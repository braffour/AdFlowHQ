# 🛠️ Vault File Setup

## 1. Configuration Directory

Create a `vault/config` folder. (This example doesn't need extra config files, but it's useful if switching to non-dev mode later.)

## 2. Persistent Data

Ensure `vault/data` exists — it's where Vault stores data if restarted. Using a file storage backend like this makes the setup more durable.

---

## ✅ Start Vault
```bash
docker-compose up -d
```

## ✅ Use Vault CLI
```bash
export VAULT_ADDR=http://127.0.0.1:8200
export VAULT_TOKEN=root
vault status
```

## ✅ Check Vault UI
[http://localhost:8200/ui/](http://localhost:8200/ui/)

---

## ✅ Store Gemini API Key & Ad Platform Secrets

```bash
vault kv put secret/adflowhq/gemini api_key="YOUR_GEMINI_API_KEY"
vault kv put secret/adflowhq/google-ads google_ads_key="YOUR_GOOGLE_ADS_KEY"
vault kv put secret/adflowhq/facebook-ads facebook_ads_key="YOUR_FACEBOOK_ADS_KEY"
vault kv put secret/adflowhq/callrail api_key="YOUR_CALLRAIL_API_KEY"
```

Or, for more granular secrets:

```bash
vault kv put secret/gemini api_key="YOUR_GEMINI_API_KEY"
vault kv put secret/google_ads client_id="YOUR_CLIENT_ID" client_secret="YOUR_CLIENT_SECRET" refresh_token="YOUR_REFRESH_TOKEN"
vault kv put secret/facebook_ads app_id="YOUR_APP_ID" app_secret="YOUR_APP_SECRET" access_token="YOUR_ACCESS_TOKEN"
vault kv put secret/callrail api_key="YOUR_CALLRAIL_API_KEY"
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
    








