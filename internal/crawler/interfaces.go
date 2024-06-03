package crawler

import (
	"sportshot/pkg/utils/models/event"
)

type Crawler interface {
	Crawl() []event.SportEvent
	SaveToMongo([]event.SportEvent)
}
