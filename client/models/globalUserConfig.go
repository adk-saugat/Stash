package models

import "encoding/json"

type GlobalUserConfig struct {
	Username  string `json:"username"`
	UserEmail string `json:"userEmail"`
}

func GlobalUserConfigFromJSON(data []byte) (*GlobalUserConfig, error) {
	var guc GlobalUserConfig
	err := json.Unmarshal(data, &guc)
	if err != nil {
		return nil, err
	}
	return &guc, nil
}

