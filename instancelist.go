package objectrocket

import (
	"net/http"
)
// InstanceService interacts with the Create instance section of the API
type InstanceService struct {
	client *Client
}
// Structure for instance creation request
type InstanceDetail struct {
	Name string `json:"name"`
	Plan int `json:"plan"`
	Service string `json:"service"`
	Type string `json:"type"`
	Version string `json:"version"`
	Zone string `json:"zone"`
}

// Create new instance
func (s *InstanceService) Create_Instance(data InstanceDetail) (InstanceDetail, *http.Response, error) {
	req, err := s.client.NewRequest("POST", "v2/instances", data)

	if err != nil {
		return InstanceDetail{}, nil, err
	}

  var inst InstanceDetail
	resp, err := s.client.Do(req, &inst)

	if err != nil {
		return InstanceDetail{}, resp, err
	}

	return inst, resp, nil
}
