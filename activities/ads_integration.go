package adflowhq.activities

import (
    "context"
    "os"
    "github.com/hashicorp/vault/api"
)

func getSecretFromVault(path, key string) (string, error) {
    vaultAddr := os.Getenv("VAULT_ADDR")
    vaultToken := os.Getenv("VAULT_TOKEN")
    client, err := api.NewClient(&api.Config{Address: vaultAddr})
    if err != nil {
        return "", err
    }
    client.SetToken(vaultToken)
    secret, err := client.Logical().Read(path)
    if err != nil || secret == nil || secret.Data == nil {
        return "", err
    }
    val, ok := secret.Data[key].(string)
    if !ok {
        return "", err
    }
    return val, nil
}

func SyncGoogleAds(ctx context.Context, leadID string) error {
    apiKey, err := getSecretFromVault("secret/adflowhq/google-ads", "google_ads_key")
    if err != nil {
        return err
    }
    // TODO: Use apiKey to call Google Ads API for leadID
    return nil
}

func SyncFacebookAds(ctx context.Context, leadID string) error {
    apiKey, err := getSecretFromVault("secret/adflowhq/facebook-ads", "facebook_ads_key")
    if err != nil {
        return err
    }
    // TODO: Use apiKey to call Facebook Ads API for leadID
    return nil
}

func SyncCallRail(ctx context.Context, leadID string) error {
    apiKey, err := getSecretFromVault("secret/adflowhq/callrail", "api_key")
    if err != nil {
        return err
    }
    // TODO: Use apiKey to call CallRail API for leadID
    return nil
}