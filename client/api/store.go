package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/adk-saugat/stash/client/models"
)

type ShareRequest struct {
	ProjectId   string       `json:"project_id"`
	ProjectName string       `json:"project_name"`
	Store       models.Store `json:"store"`
}

type ShareResponse struct {
	Message   string `json:"message"`
	StoreId   string `json:"store_id"`
	ProjectId string `json:"project_id"`
	Error     string `json:"error"`
}

func ShareStore(token string, projectId, projectName string, store models.Store) (*ShareResponse, error) {
	payload := ShareRequest{
		ProjectId:   projectId,
		ProjectName: projectName,
		Store:       store,
	}
	jsonData, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", ServerURL+"/api/share", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("could not create request")
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not connect to server")
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var shareResp ShareResponse
	json.Unmarshal(body, &shareResp)

	if resp.StatusCode == http.StatusUnauthorized {
		return nil, fmt.Errorf("unauthorized: please login first")
	}
	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("share failed: %s", shareResp.Error)
	}

	return &shareResp, nil
}
