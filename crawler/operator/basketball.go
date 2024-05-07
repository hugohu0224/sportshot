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
	doc := bson.M{"date": time.Now().Format("2020-01-01"), "events": events}

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
	// initial Colly
	c := colly.NewCollector()

	// to avoid return before the subsequent Visit is completed,
	// created a channel to receive the result (blocking).
	resultChan := make(chan []event.SportEvent, 1)

	// crawl logic of main
	c.OnHTML("#tbl_inplay > tbody", func(e *colly.HTMLElement) {
		currentTimestamp := time.Now().Unix()
		var events []event.SportEvent
		zap.S().Info("crawling basketball events ")
		// // find the "tr" list (rows)
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			// Iterate over "td" in "tr" (events in row)
			ev := event.SportEvent{}
			columnIdx := 0
			el.ForEach("td", func(index int, td *colly.HTMLElement) {
				// since many different columns use the same class tag,
				// the event is retrieved in sequential order.
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
			//  collecting events
			events = append(events, ev)
		})
		resultChan <- events
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL.String())
	})

	if err := c.Visit("https://tw.betsapi.com/ciz/basketball"); err != nil {
		return nil
	}

	return <-resultChan
}
