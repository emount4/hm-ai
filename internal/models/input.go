package models

type WorkerInput struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Skills   string  `json:"skills"`
	District string  `json:"district"`
	City     string  `json:"city"`
	Rating   float32 `json:"rating"`
}

type Request struct {
	Workers []WorkerInput `json:"workers"`
	Task    string        `json:"task"`
	TopK    int           `json:"top_k"`
}
