package crawler

import (
	"sportshot/pkg/utils/models/events"
)

type Crawler interface {
	Crawl() []events.SportEvent
	SaveToMongo([]events.SportEvent)
}
