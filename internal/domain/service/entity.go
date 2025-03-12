package service

type Service struct {
	ID          uint   `json:"serviceId"`
	Name        string `json:"name"`
	Version     string `json:"version"`
	AddedBy     string `json:"addedBy"`
	ReleaseDate string `json:"releaseDate"`
	Assigned    *bool  `json:"assigned"`
	Status      int    `json:"status"`
}
