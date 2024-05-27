package event

type SportEvent struct {
	LeagueName string `bson:"leagueName" json:"leagueName"`
	RaceTime   string `bson:"raceTime"   json:"raceTime"`
	HomeName   string `bson:"homeName"   json:"homeName"`
	Score      string `bson:"score"      json:"score"`
	AwayName   string `bson:"awayName"   json:"awayName"`
	HomeOdds   string `bson:"homeOdds"   json:"homeOdds"`
	AwayOdds   string `bson:"awayOdds"   json:"awayOdds"`
	Date       string `bson:"date"       json:"date"`
	Timestamp  int    `bson:"timestamp"  json:"timestamp"`
}

type SearchEventsForm struct {
	LeagueName string `json:"leagueName" form:"leagueName"`
	HomeName   string `json:"homeName" form:"homeName"`
	AwayName   string `json:"awayName" form:"awayName"`
	SportType  string `json:"sportType" form:"sportType"`
	StartDate  string `json:"startDate" form:"startDate"`
	EndDate    string `json:"endDate" form:"endDate"`
}

func (g *SearchEventsForm) SetDefaults() {

}
