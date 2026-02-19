package models

type Recommendation struct {
	WorkerID uint    `json:"worker_id"`
	Name     string  `json:"name"`
	Score    float32 `json:"similarity_score"`
	Rating   float32 `json:"rating"`
}

type Response struct {
	Recommendations []Recommendation `json:"recommendations"`
}
