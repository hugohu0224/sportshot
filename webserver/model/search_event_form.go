package model

type SearchEventsForm struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Date string `json:"date"`
}

// SetDefaults 設置結構體的預設值
func (g *SearchEventsForm) SetDefaults() {

}
