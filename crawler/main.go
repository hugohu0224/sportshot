package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly/v2"
	"go.uber.org/zap"
	"log"
	"os"
	"sportshot/crawler/model"
	"strings"
	"time"
)

func main() {
	// initialize logger
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)

	c := colly.NewCollector()
	c.OnHTML("#tbl_inplay > tbody", func(e *colly.HTMLElement) {
		zap.S().Info("OnHTML")
		currentTimestamp := time.Now().Unix()
		var events []model.SportInfo
		// 找到tr (row)
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			event := model.SportInfo{}
			// 因多個不同欄位使用同個class tag，因此使用順序方式獲取event
			columnIdx := 0
			// 遍歷tr中的td (event in row)
			el.ForEach("td", func(index int, td *colly.HTMLElement) {
				event.Timestamp = int(currentTimestamp)
				switch columnIdx {
				case 0:
					event.LeagueName = strings.TrimSpace(td.Text)
				case 1:
					event.RaceTime = td.ChildText("span.race-time")
				case 2:
					event.HomeName = strings.TrimSpace(td.Text)
				case 3:
					event.Score = strings.TrimSpace(td.Text)
				case 4:
					event.AwayName = strings.TrimSpace(td.Text)
				case 5:
					event.HomeOdds = td.Text
				case 6:
					event.AwayOdds = td.Text
				}
				columnIdx += 1
			})
			// 美化output
			jsonData, err := json.MarshalIndent(events, "", "    ")
			if err != nil {
				log.Println("Error marshaling data:", err)
				return
			}
			fmt.Println("Data extracted:\n", string(jsonData))
			events = append(events, event)
		})

		file, err := os.Create("events.json")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		encoder := json.NewEncoder(file)
		err = encoder.Encode(events)
		if err != nil {
			panic(err)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://tw.betsapi.com/ciz/basketball")
}
