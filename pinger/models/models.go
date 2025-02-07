package models

type PingResult struct {
	IP       string `json:"ip"`
	PingTime string `json:"ping_time"`
	Date     string `json:"date"`
}

type PingResults struct {
	Results []PingResult `json:"ping_results"`
}
