package entity

type SimpleDemo struct {
	Name       string `json:"name"`
	StuffNo    string `json:"stuff_no"`
	Department string `json:"department"`
	Password   string `json:"password"`
	Active     bool   `json:"active"`
	Info       string `json:"info"`
}
