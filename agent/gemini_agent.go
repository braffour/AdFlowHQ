package adflowhq.agent

import (
	"context"
	"io/ioutil"
	"net/http"
	"os"
)

func GeminiAgent(ctx context.Context, prompt string) (string, error) {
    apiKey := os.Getenv("GEMINI_API_KEY")
    req, _ := http.NewRequest("POST", "https://api.gemini.example.com/v1/chat", nil)
    req.Header.Set("Authorization", "Bearer " + apiKey)
    // TODO: set JSON body with prompt
    client := &http.Client{}
    resp, err := client.Do(req)
    if err!=nil { return "", err }
    defer resp.Body.Close()
    b, _ := ioutil.ReadAll(resp.Body)
    return string(b), nil
}