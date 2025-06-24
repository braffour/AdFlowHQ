package activities

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/vault/api"
)

func getSecretFromVault(path, key string) (string, error) {
	vaultAddr := os.Getenv("VAULT_ADDR")
	if vaultAddr == "" {
		vaultAddr = "http://localhost:8200"
	}

	vaultToken := os.Getenv("VAULT_TOKEN")
	if vaultToken == "" {
		return "", fmt.Errorf("VAULT_TOKEN environment variable not set")
	}

	client, err := api.NewClient(&api.Config{Address: vaultAddr})
	if err != nil {
		return "", fmt.Errorf("failed to create vault client: %w", err)
	}
	client.SetToken(vaultToken)

	secret, err := client.Logical().Read(path)
	if err != nil {
		return "", fmt.Errorf("failed to read secret from vault: %w", err)
	}
	if secret == nil || secret.Data == nil {
		return "", fmt.Errorf("no secret data found at path: %s", path)
	}

	val, ok := secret.Data[key].(string)
	if !ok {
		return "", fmt.Errorf("secret key %s not found or not a string", key)
	}
	return val, nil
}

func SyncGoogleAds(ctx context.Context, leadID string) error {
	log.Printf("Syncing Google Ads for lead: %s", leadID)

	apiKey, err := getSecretFromVault("secret/adflowhq/google-ads", "google_ads_key")
	if err != nil {
		return fmt.Errorf("failed to get Google Ads API key: %w", err)
	}

	// TODO: Implement actual Google Ads API integration
	// For now, just log the action
	log.Printf("Would sync lead %s to Google Ads with API key: %s...", leadID, apiKey[:8])

	return nil
}

func SyncFacebookAds(ctx context.Context, leadID string) error {
	log.Printf("Syncing Facebook Ads for lead: %s", leadID)

	apiKey, err := getSecretFromVault("secret/adflowhq/facebook-ads", "facebook_ads_key")
	if err != nil {
		return fmt.Errorf("failed to get Facebook Ads API key: %w", err)
	}

	// TODO: Implement actual Facebook Ads API integration
	// For now, just log the action
	log.Printf("Would sync lead %s to Facebook Ads with API key: %s...", leadID, apiKey[:8])

	return nil
}

func SyncCallRail(ctx context.Context, leadID string) error {
	log.Printf("Syncing CallRail for lead: %s", leadID)

	apiKey, err := getSecretFromVault("secret/adflowhq/callrail", "api_key")
	if err != nil {
		return fmt.Errorf("failed to get CallRail API key: %w", err)
	}

	// TODO: Implement actual CallRail API integration
	// For now, just log the action
	log.Printf("Would sync lead %s to CallRail with API key: %s...", leadID, apiKey[:8])

	return nil
}