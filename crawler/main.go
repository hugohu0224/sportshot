package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly/v2"
	"log"
	"sportshot/crawler/model"
	"strings"
	"time"
)

func main() {
	c := colly.NewCollector()
	c.OnHTML("#tbl_inplay > tbody", func(e *colly.HTMLElement) {
		currentTimestamp := time.Now().Unix()
		// 找到tr (row)
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			data := model.SportInfo{}
			// 因多個不同欄位使用同個class tag，因此使用順序方式獲取data
			columnIdx := 0
			// 遍歷tr中的td (data in row)
			el.ForEach("td", func(index int, td *colly.HTMLElement) {
				data.Timestamp = int(currentTimestamp)
				switch columnIdx {
				case 0:
					data.LeagueName = strings.TrimSpace(td.Text)
				case 1:
					data.RaceTime = td.ChildText("span.race-time")
				case 2:
					data.HomeName = strings.TrimSpace(td.Text)
				case 3:
					data.Score = strings.TrimSpace(td.Text)
				case 4:
					data.AwayName = strings.TrimSpace(td.Text)
				case 5:
					data.HomeOdds = td.Text
				case 6:
					data.AwayOdds = td.Text
				}
				columnIdx += 1
			})
			// 美化output
			jsonData, err := json.MarshalIndent(data, "", "    ")
			if err != nil {
				log.Println("Error marshaling data:", err)
				return
			}
			fmt.Println("Data extracted:\n", string(jsonData))
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://tw.betsapi.com/ciz/basketball")
}
