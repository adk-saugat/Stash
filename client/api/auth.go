package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const ServerURL = "http://localhost:8080"

var ErrUserNotFound = errors.New("user not found")

type AuthResponse struct {
	Message string `json:"message"`
	Email   string `json:"email"`
	Token   string `json:"token"`
	Error   string `json:"error"`
}

func Login(email, password string) (*AuthResponse, error) {
	payload := map[string]string{"email": email, "password": password}
	jsonData, _ := json.Marshal(payload)

	resp, err := http.Post(ServerURL+"/api/login", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("could not connect to server")
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var authResp AuthResponse
	json.Unmarshal(body, &authResp)

	if resp.StatusCode == http.StatusNotFound {
		return nil, ErrUserNotFound
	}
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, fmt.Errorf("incorrect password")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("login failed: %s", authResp.Error)
	}

	return &authResp, nil
}

func Register(username, email, password string) (*AuthResponse, error) {
	payload := map[string]string{"username": username, "email": email, "password": password}
	jsonData, _ := json.Marshal(payload)

	resp, err := http.Post(ServerURL+"/api/register", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("could not connect to server")
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var authResp AuthResponse
	json.Unmarshal(body, &authResp)

	if resp.StatusCode == http.StatusConflict {
		return nil, fmt.Errorf("user already exists")
	}
	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("registration failed: %s", authResp.Error)
	}

	return &authResp, nil
}
