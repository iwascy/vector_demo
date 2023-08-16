package azure

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Configuration struct {
	ResourceName    string
	DeploymentName  string
	APIKey          string
	APIVersion      string
	EndpointPattern string
}

type DocumentInput struct {
	Input string `json:"input"`
}

type EmbeddingResponse struct {
	Object string     `json:"object"`
	Data   []Embed    `json:"data"`
	Model  string     `json:"model"`
	Usage  UsageStats `json:"usage"`
}

type Embed struct {
	Object    string    `json:"object"`
	Index     int       `json:"index"`
	Embedding []float64 `json:"embedding"`
}

type UsageStats struct {
	PromptTokens int `json:"prompt_tokens"`
	TotalTokens  int `json:"total_tokens"`
}

func NewDefaultConfiguration(resourceName, deploymentName, key string) Configuration {
	return Configuration{
		APIVersion:      "2023-05-15",
		EndpointPattern: "https://%s.openai.azure.com/openai/deployments/%s/embeddings?api-version=%s",
		ResourceName:    resourceName,
		DeploymentName:  deploymentName,
		APIKey:          key,
	}
}

func (c *Configuration) GetEmbedding(input DocumentInput) (*EmbeddingResponse, error) {
	url := fmt.Sprintf(c.EndpointPattern, c.ResourceName, c.DeploymentName, c.APIVersion)

	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("api-key", c.APIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var embeddingResponse EmbeddingResponse
	err = json.Unmarshal(body, &embeddingResponse)
	if err != nil {
		return nil, err
	}

	return &embeddingResponse, nil
}
