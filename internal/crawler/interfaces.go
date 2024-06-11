package crawler

import (
	"sportshot/pkg/utils/models"
)

type Crawler interface {
	Crawl() []models.SportEvent
	SaveToMongo([]models.SportEvent)
}
