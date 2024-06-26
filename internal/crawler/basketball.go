package crawler

import (
	"context"
	"fmt"
	"github.com/gocolly/colly/v2"
	"go.uber.org/zap"
	"sportshot/pkg/utils/global"
	"sportshot/pkg/utils/models"
	"strings"
	"time"
)

type BasketballCrawler struct {
}

func (bc *BasketballCrawler) Crawl() []models.SportEvent {
	// initial Colly
	c := colly.NewCollector()

	// to avoid return before the subsequent Visit is completed,
	// created a channel to receive the result (blocking).
	resultChan := make(chan []models.SportEvent, 1)

	// crawl logic
	c.OnHTML("#tbl_inplay > tbody", func(e *colly.HTMLElement) {
		currentTimestamp := time.Now().Unix()
		var eventList []models.SportEvent
		zap.S().Info("crawling basketball events ")
		// find the "tr" (rows)
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			// "td" in "tr" (single row)
			ev := models.SportEvent{}
			columnIdx := 0
			el.ForEach("td", func(index int, td *colly.HTMLElement) {
				// since many different columns use the same class tag,
				// the events is retrieved in sequential order.
				ev.Timestamp = int(currentTimestamp)
				ev.Date = time.Now().Format("2006-01-02")
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
			eventList = append(eventList, ev)
		})
		resultChan <- eventList
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL.String())
	})

	if err := c.Visit("https://betsapi.com/ciz/basketball"); err != nil {
		return nil
	}

	return <-resultChan
}

func (bc *BasketballCrawler) SaveToMongo(events []models.SportEvent) {
	// connect to mongo
	databaseName := "sportevents"
	collectionName := "basketball"
	collection := global.MongodbClient.Database(databaseName).Collection(collectionName)

	// convert SportEvent type to interface{} due to mongodb insertion requirements
	var docs []interface{}
	for _, e := range events {
		docs = append(docs, e)
	}

	// start to insert
	zap.S().Infof("start to insert data to Mongodb.%s.%s", databaseName, collectionName)
	result, err := collection.InsertMany(context.TODO(), docs)
	if err != nil {
		zap.S().Errorf(err.Error())
	}
	zap.S().Infof("inserted docs %s to Mongodb.%s.%s", result.InsertedIDs, databaseName, collectionName)

}
