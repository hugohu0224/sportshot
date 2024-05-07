package operator

import (
	"context"
	"fmt"
	"github.com/gocolly/colly/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
	"sportshot/crawler/global"
	"sportshot/utils/models/event"
	"strings"
	"time"
)

type BasketballCrawler struct {
}

func (cr *BasketballCrawler) SaveToMongo(events []event.SportEvent) {
	// reorg data
	doc := bson.M{"date": time.Now().Format("2006-01-02"), "events": events}

	// connect to mongo
	databaseName := "sportevents"
	collectionName := "basketball"
	collection := global.MongodbClient.Database(databaseName).Collection(collectionName)
	zap.S().Infof("start to insert data to Mongodb.%s.%s", databaseName, collectionName)

	// insert data
	if _, err := collection.InsertOne(context.TODO(), doc); err != nil {
		zap.S().Error("failed to insert document:", err)
	} else {
		zap.S().Infof("Mongodb.%s.%s data inserted", databaseName, collectionName)
	}
}

func (cr *BasketballCrawler) Crawl() []event.SportEvent {
	c := colly.NewCollector()
	// 避免後續Visit尚未完成就return，建立一個通道來接收result(阻塞)
	resultChan := make(chan []event.SportEvent, 1)

	c.OnHTML("#tbl_inplay > tbody", func(e *colly.HTMLElement) {
		currentTimestamp := time.Now().Unix()
		var events []event.SportEvent
		zap.S().Info("crawling basketball events ")
		// 找到tr列表 (rows)
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			// 遍歷tr中的td (events in row)
			ev := event.SportEvent{}
			columnIdx := 0

			el.ForEach("td", func(index int, td *colly.HTMLElement) {
				// 因多個不同欄位使用同個class tag，因此使用順序方式獲取event
				ev.Timestamp = int(currentTimestamp)
				switch columnIdx {
				case 0:
					ev.LeagueName = strings.TrimSpace(td.Text)
				case 1:
					ev.RaceTime = td.ChildText("span.race-time")
				case 2:
					ev.HomeName = strings.TrimSpace(td.Text)
				case 3:
					ev.Score = strings.TrimSpace(td.Text)
				case 4:
					ev.AwayName = strings.TrimSpace(td.Text)
				case 5:
					ev.HomeOdds = td.Text
				case 6:
					ev.AwayOdds = td.Text
				}
				columnIdx += 1
			})

			// 收集events
			events = append(events, ev)
		})

		resultChan <- events
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL.String())
	})

	err := c.Visit("https://tw.betsapi.com/ciz/basketball")

	if err != nil {
		return nil
	}

	return <-resultChan
}
