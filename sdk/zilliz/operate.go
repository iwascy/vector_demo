package zilliz

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ZillizSDK struct to hold SDK configuration
type ZillizSDK struct {
	Endpoint   string
	AuthToken  string
	httpClient *http.Client
}

// New initializes a new instance of ZillizSDK
func New(endpoint, authToken string) *ZillizSDK {
	return &ZillizSDK{
		Endpoint:   endpoint,
		AuthToken:  authToken,
		httpClient: &http.Client{},
	}
}

// VectorData represents the data structure for inserting vectors
type VectorData struct {
	CollectionName string `json:"collectionName"`
	Data           Data   `json:"data"`
}

type Data struct {
	Vector []float64 `json:"vector"`
	LiveId string    `json:"liveId"`
	Text   string    `json:"text"`
}

type InsertResp struct {
	Code int `json:"code"`
	Data struct {
		InsertCount int      `json:"insertCount"`
		InsertIds   []string `json:"insertIds"`
	} `json:"data"`
}

// InsertCollection inserts vector into a Zilliz collection
func (sdk *ZillizSDK) InsertCollection(data VectorData) (resp *InsertResp, err error) {
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal data: %v", err)
	}

	url := sdk.Endpoint + "/vector/insert"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Add("Authorization", "Bearer "+sdk.AuthToken)
	res, err := sdk.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var insertResp = &InsertResp{}
	err = json.Unmarshal(body, insertResp)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %v", err)
	}

	return insertResp, nil
}

type QueryReq struct {
	CollectionName string    `json:"collectionName"`
	Filter         string    `json:"filter"`
	Vector         []float64 `json:"vector"`
}

type QueryResp struct {
	Code int           `json:"code"`
	Data []QueryResult `json:"data"`
}

type QueryData struct {
	QueryResult
}

type QueryResult struct {
	Id     int       `json:"id"`
	Vector []float64 `json:"vector"`
	Text   string    `json:"text"`
	LiveId string    `json:"liveId"`
}

func (sdk *ZillizSDK) SearchCollection(query QueryReq) (resp *QueryResp, err error) {
	payload, err := json.Marshal(query)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal query: %v", err)
	}

	url := sdk.Endpoint + "/vector/search"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Add("Authorization", "Bearer "+sdk.AuthToken)
	req.Header.Add("Accept", "application/json")
	res, err := sdk.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var queryResp = &QueryResp{}
	err = json.Unmarshal(body, queryResp)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %v", err)
	}

	fmt.Printf("queryResp: %+v \n", string(body))

	return queryResp, nil
}
