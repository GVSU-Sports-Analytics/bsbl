package bsbl

type Request struct {
	ApiKey    string `json:"api_key"`
	Team1Name string `json:"team_1_name"`
	Team1Yr   string `json:"team_1_yr"`
	Team2Name string `json:"team_2_name"`
	Team2Yr   string `json:"team_2_yr"`
}
