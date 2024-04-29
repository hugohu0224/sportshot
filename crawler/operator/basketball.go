package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"sportshot/crawler/model"
	"strings"
	"time"
)

type basketballCrawler struct {
}

func (cr *basketballCrawler) crawl(url string) []model.SportEvent {
	c := colly.NewCollector()
	// 避免後續Visit尚未完成就return，建立一個通道來接收result(阻塞)
	resultChan := make(chan []model.SportEvent, 1)

	c.OnHTML("#tbl_inplay > tbody", func(e *colly.HTMLElement) {
		currentTimestamp := time.Now().Unix()
		var events []model.SportEvent

		// 找到tr列表 (rows)
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			// 遍歷tr中的td (events in row)
			event := model.SportEvent{}
			columnIdx := 0

			el.ForEach("td", func(index int, td *colly.HTMLElement) {
				// 因多個不同欄位使用同個class tag，因此使用順序方式獲取event
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

			//// 美化stdout
			//jsonData, err := json.MarshalIndent(events, "", "    ")
			//if err != nil {
			//	log.Println("Error marshaling data:", err)
			//	return
			//}
			//fmt.Println("Data extracted:\n", string(jsonData))

			// 收集events
			events = append(events, event)
		})

		// localize data
		//file, err := os.Create("events.json")
		//if err != nil {
		//	panic(err)
		//}
		//defer file.Close()
		//encoder := json.NewEncoder(file)
		//err = encoder.Encode(events)
		//if err != nil {
		//	panic(err)
		//}
		resultChan <- events
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	//err := c.Visit("https://tw.betsapi.com/ciz/basketball")
	err := c.Visit(url)

	if err != nil {
		return nil
	}

	return <-resultChan
}
