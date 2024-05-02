package event

type SportEvent struct {
	LeagueName string `bson:"leagueName" json:"leagueName"`
	RaceTime   string `bson:"raceTime" json:"raceTime"`
	HomeName   string `bson:"homeName" json:"homeName"`
	Score      string `bson:"score" json:"score"`
	AwayName   string `bson:"awayName" json:"awayName"`
	HomeOdds   string `bson:"homeOdds" json:"homeOdds"`
	AwayOdds   string `bson:"awayOdds" json:"awayOdds"`
	Timestamp  int    `bson:"timestamp" json:"timestamp"`
}

type SearchEventsForm struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Date string `json:"date"`
}

// SetDefaults 設置結構體的預設值
func (g *SearchEventsForm) SetDefaults() {

}
