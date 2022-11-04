package main

// Define Player type
type Player struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Team       string `json:"team"`
	Position   string `json:"position"`
	Rosterable bool   `json:"rosterable"`
}

func main() {
}
