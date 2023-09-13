package storage

import (
	"encoding/json"
	"fmt"
)

type SaiStorageSaveRequest struct {
	Collection string      `json:"collection"`
	Data       interface{} `json:"data"`
	Options    interface{} `json:"options,omitempty"`
}

type SaiStorageUpdateRequest struct {
	Collection string      `json:"collection"`
	Select     interface{} `json:"select"`
	Data       interface{} `json:"data"`
	Options    interface{} `json:"options,omitempty"`
}

type SaiStorageUpsertRequest = SaiStorageUpdateRequest

type SaiStorageRemoveRequest struct {
	Collection string      `json:"collection"`
	Select     interface{} `json:"select"`
	Options    interface{} `json:"options,omitempty"`
}

type SaiStorageChangeResponse struct {
	Status string `json:"Status"`
	Result string `json:"Result"`
}

func (saiStorage *SaiStorage) Save(request SaiStorageSaveRequest) (*SaiStorageChangeResponse, error) {
	response, err := saiStorage.makeChangeRequest("save", request)
	if err != nil {
		return nil, err
	}

	// Return the parsed results
	return response, nil
}

func (saiStorage *SaiStorage) Update(request SaiStorageUpdateRequest) (*SaiStorageChangeResponse, error) {
	response, err := saiStorage.makeChangeRequest("update", request)
	if err != nil {
		return nil, err
	}

	// Return the parsed results
	return response, nil
}

func (saiStorage *SaiStorage) Upsert(request SaiStorageUpsertRequest) (*SaiStorageChangeResponse, error) {
	response, err := saiStorage.makeChangeRequest("upsert", request)
	if err != nil {
		return nil, err
	}

	// Return the parsed results
	return response, nil
}

func (saiStorage *SaiStorage) Remove(request SaiStorageRemoveRequest) (*SaiStorageChangeResponse, error) {
	response, err := saiStorage.makeChangeRequest("remove", request)
	if err != nil {
		return nil, err
	}

	// Return the parsed results
	return response, nil
}

func (saiStorage *SaiStorage) makeChangeRequest(method string, request interface{}) (*SaiStorageChangeResponse, error) {

	requestBody, err := json.Marshal(request)

	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err)
	}

	// Make the request
	response, err := saiStorage.makeRequest(method, requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}
	defer response.Body.Close()

	// Parse the response body into the struct
	var result SaiStorageChangeResponse
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response body: %v", err)
	}

	// Return the parsed results
	return &result, nil
}
